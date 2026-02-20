package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/adapter/db"
	workerAdapter "github.com/kritpi/499-senior-project-trip-service/internal/adapter/worker"
	workerSvc "github.com/kritpi/499-senior-project-trip-service/internal/adapter/worker/service"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/service"
	handler "github.com/kritpi/499-senior-project-trip-service/internal/handler/rest-api"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/socket"
	"github.com/kritpi/499-senior-project-trip-service/internal/middleware"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository"
	redisrepository "github.com/kritpi/499-senior-project-trip-service/internal/repository/redis-repository"
	"github.com/kritpi/499-senior-project-trip-service/internal/router"
	"github.com/kritpi/499-senior-project-trip-service/property"
)

func main() {
	// Load config
	cfg, err := property.LoadConfig("./config", "local")
	if err != nil {
		log.Fatalf("error loading config: %+v", err)
	}

	// --- Fiber (REST) ---
	app := fiber.New()
	middleware.SetupCORS(app, cfg)

	// init db
	ctx := context.Background()
	// db
	pool := db.NewPostgres(ctx, cfg.Database.Url)
	defer pool.Close()

	redisCli := db.NewRedisClient(ctx, *cfg)
	defer redisCli.Close()

	// repo
	repo := repository.New(ctx, pool, *cfg)
	redisRepo := redisrepository.New(ctx, redisCli, *cfg)
	svc := service.New(&repo, &redisRepo, cfg)
	h := handler.NewRestApi(svc, *cfg)

	router.SetupRouter(app, h, *cfg)

	// --- Socket.IO ---
	// Create server first (with nil handler temporarily)
	socketServer := socket.NewSocket(nil, *cfg)

	// Create handler with server instance
	socketHandler := socket.NewSocketIO(svc, socketServer)

	// Register handlers to server
	socket.RegisterHandlers(socketServer, socketHandler)

	// --- Workers ---
	// Initialize workers to run in the same process as the server
	upsertWorker := workerSvc.NewActivityUpsertWorker(svc, redisCli)
	broadcastWorker := workerSvc.NewActivityBroadcastWorker(svc, redisCli, socketServer)
	worker := workerAdapter.NewWorker(redisCli, upsertWorker, broadcastWorker)

	// Start workers in a goroutine
	go func() {
		log.Println("Starting workers...")
		if err := worker.Run(context.Background()); err != nil {
			log.Fatalf("Worker error: %+v", err)
		}
	}()

	// --- HTTP multiplexer ---
	mux := http.NewServeMux()

	// REST → /api/*
	mux.Handle("/api/", adaptor.FiberApp(app))

	// Socket.IO → /socket.io/*
	router.RegisterSocket(mux, socketServer, socketHandler, *cfg)

	port := cfg.Server.Port
	log.Printf("Starting server on port: %s", port)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
