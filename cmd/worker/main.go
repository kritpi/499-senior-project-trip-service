package main

import (
	"context"
	"log"

	"github.com/kritpi/499-senior-project-trip-service/internal/adapter/db"
	workerAdapter "github.com/kritpi/499-senior-project-trip-service/internal/adapter/worker"
	workerSvc "github.com/kritpi/499-senior-project-trip-service/internal/adapter/worker/service"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/service"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository"
	redisrepository "github.com/kritpi/499-senior-project-trip-service/internal/repository/redis-repository"
	"github.com/kritpi/499-senior-project-trip-service/property"
)

func main() {
	cfg, err := property.LoadConfig("./config", "local")
	if err != nil {
		log.Fatalf("error loading worker config: %+v", err)
	}

	ctx := context.Background()
	postgres := db.NewPostgres(ctx, cfg.Database.Url)
	defer postgres.Close()

	redis := db.NewRedisClient(ctx, *cfg)
	defer redis.Close()
	repo := repository.New(ctx, postgres, *cfg)
	redisRepo := redisrepository.New(ctx, redis, *cfg)

	coreSvc := service.New(&repo, &redisRepo, cfg)
	upsertWorker := workerSvc.NewActivityUpsertWorker(coreSvc, redis)
	// Note: broadcast worker needs socket server which is in server process
	// For now, we'll pass nil and handle this architectural issue separately
	broadcastWorker := workerSvc.NewActivityBroadcastWorker(coreSvc, redis, nil)

	worker := workerAdapter.NewWorker(redis, upsertWorker, broadcastWorker)

	log.Printf("Starting worker...")

	// Run the worker - this should block indefinitely
	if err := worker.Run(context.Background()); err != nil {
		log.Fatalf("Worker error: %+v", err)
	}
}
