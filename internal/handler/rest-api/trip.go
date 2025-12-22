package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
)

func (h *restHandler) GetTrips(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.Trip

	err := c.QueryParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	trips, err := h.svc.GetTrips(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to fetch trips",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"trips": trips,
	})
}
