package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/kritpi/499-senior-project-trip-service/internal/handler/rest-api"
	"github.com/kritpi/499-senior-project-trip-service/internal/middleware"
	"github.com/kritpi/499-senior-project-trip-service/property"
)

func SetupRouter(app *fiber.App, h handler.RestHandler, cfg property.Property) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Auth
	v1.Post("/auth/google", h.GoogleAuth)

	// Trip
	trip := v1.Group("trip")
	trip.Use(middleware.AuthMiddleware(cfg))

	trip.Put("/", h.UpsertTrip)
	trip.Get("/", h.GetMemberTrips)
	// Get trips (with member id) => my trip, my involved trip
	// Get trip by id (id)

}
