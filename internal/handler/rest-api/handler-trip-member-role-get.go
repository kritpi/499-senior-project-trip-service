package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

func (h *restHandler) GetTripMemberRole(c *fiber.Ctx) error {
	ctx := c.Context()

	jwtClaims := c.Locals("users").(*model.JWTCustomClaims)

	tripIdParam := c.Params("id")
	tripId, err := strconv.Atoi(tripIdParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid trip ID",
		})
	}

	req := dto.GetTripMemberRoleRequest{
		TripId: tripId,
	}

	resp, err := h.svc.GetTripMemberRole(ctx, *req.ToDomain(jwtClaims.ID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetTripMemberRoleResponse{}.FromDomain(*resp))

}
