package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/kritpi/499-senior-project-trip-service/internal/handler/rest-api"
)

func SetupRouter(app *fiber.App, h handler.RestHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/trips",h.GetTrips)
	
}
