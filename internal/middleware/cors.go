package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kritpi/499-senior-project-trip-service/property"
)

// SetupCORS configures CORS middleware for the application
func SetupCORS(app *fiber.App, cfg *property.Property) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:5173, http://localhost:4200",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS",
		MaxAge:       3600,
	}))
}
