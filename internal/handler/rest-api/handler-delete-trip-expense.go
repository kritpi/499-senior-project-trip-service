package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

func (h *restHandler) DeleteTripExpense(c *fiber.Ctx) error {
	ctx := c.Context()

	jwtClaims := c.Locals("user").(*model.JWTCustomClaims)
	tripIdParam := c.Params("id")
	tripId, err := strconv.Atoi(tripIdParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid trip ID",
		})
	}

	var deleteExpense dto.DeleteTripExpenseRequest
	deleteExpense.TripId = tripId

	if err := c.BodyParser(&deleteExpense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	err = h.svc.DeleteExpense(ctx, *deleteExpense.ToDomain(jwtClaims.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "deleted",
	})

}
