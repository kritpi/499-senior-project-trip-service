package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/adapter/db"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/service"
	handler "github.com/kritpi/499-senior-project-trip-service/internal/handler/rest-api"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/socket"
	"github.com/kritpi/499-senior-project-trip-service/internal/middleware"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository"
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

	// init deps
	ctx := context.Background()
	// db
	pool := db.NewPostgres(ctx, cfg.Database.Url)
	defer pool.Close()

	// repo
	repo := repository.New(ctx, pool, *cfg)
	svc := service.New(&repo, cfg)
	h := handler.NewRestApi(svc)

	router.SetupRouter(app, h)

	// --- Socket.IO ---
	socketHandler := socket.NewSocketIO(svc)
	socketServer := socket.NewSocket(socketHandler)

	// --- HTTP multiplexer ---
	mux := http.NewServeMux()

	// REST → /api/*
	mux.Handle("/api/", adaptor.FiberApp(app))

	// Socket.IO → /socket.io/*
	router.RegisterSocket(mux, socketServer, socketHandler)

	port := cfg.Server.Port
	log.Printf("Starting server on port: %s", port)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
