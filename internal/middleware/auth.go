package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/property"
	"github.com/kritpi/499-senior-project-trip-service/shared/utils"
)

func AuthMiddleware(cfg property.Property) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "missing authorization header",
			})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid authorization header format",
			}) 
		}

		tokenString := parts[1]

		claims, err := utils.ParseAndValidateToken(tokenString, cfg)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid or expired token",
			})
		}

		c.Locals("user", claims)

		return c.Next()
	}
}