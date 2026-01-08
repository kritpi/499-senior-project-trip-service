package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

func (h *restHandler) GetTripById(c *fiber.Ctx) error {
	ctx := c.Context()
	jwtClaims := c.Locals("user").(*model.JWTCustomClaims)

	// Get trip ID from URL parameter
	tripIdParam := c.Params("id")
	tripId, err := strconv.Atoi(tripIdParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid trip ID",
		})
	}

	// Call service
	req := domain.GetTripByIdRequest{
		TripId:   tripId,
		MemberId: jwtClaims.ID,
	}

	resp, err := h.svc.GetTripById(ctx, req)
	if err != nil {
		// Check if it's a forbidden error
		if err.Error() == "forbidden: member not part of this trip" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "you are not authorized to view this trip",
			})
		}
		// Check if it's a not found error
		if err.Error() == "trip not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "trip not found",
			})
		}
		// Other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Convert to DTO and return
	return c.Status(fiber.StatusOK).JSON(dto.GetTripByIdResponse{}.FromDomain(resp))
}
