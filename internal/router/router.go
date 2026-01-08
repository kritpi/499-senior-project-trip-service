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
	trip.Put("/", h.UpsertTrip) // Create/Update trip
	trip.Get("/", h.GetMemberTrips) // Get member's trips
	trip.Get("/:id", h.GetTripById) // Get trip by ID
	trip.Post("/invitation", h.InviteMember) // Invite member to trip
}
