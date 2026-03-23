package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/property"
	"github.com/kritpi/499-senior-project-trip-service/shared/utils"
)

func AuthMiddleware(cfg property.Property) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, err := utils.AuthenticateFromHeader(c.Get("Authorization"), cfg)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		c.Locals("user", claims)
		return c.Next()
	}

}
