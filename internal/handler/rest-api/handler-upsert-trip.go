package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

func (h *restHandler) UpsertTrip(c *fiber.Ctx) error {
	ctx := c.Context()
	jwtClaims := c.Locals("user").(*model.JWTCustomClaims)
	var req dto.UpsertTripRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	resp, err := h.svc.UpsertTrip(ctx, *req.ToDomain(jwtClaims.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.UpsertTripResponse{
		TripId: resp.TripId,
	})

}
