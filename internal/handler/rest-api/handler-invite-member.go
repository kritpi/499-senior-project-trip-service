package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

func (h *restHandler) InviteMember(c *fiber.Ctx) error {
	ctx := c.Context()
	jwtClaims := c.Locals("user").(*model.JWTCustomClaims)

	var req dto.TripInvitationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	resp, err := h.svc.InviteMember(ctx, *req.ToDomain(jwtClaims.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if resp != nil && err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "unable to invite member",
			"members": resp.Member,
		})
	}

	if resp != nil && err == nil {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "unable to invite these members",
		"members": resp.Member,
	})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "invitation sent",
	})
}
