package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
)

func (h *restHandler) GoogleAuth(c *fiber.Ctx) error {
	ctx := c.Context()
	var googleIdToken dto.GoogleIdToken

	err := c.BodyParser(&googleIdToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	tokenResp, err := h.svc.GoogleAuth(ctx, *googleIdToken.ToDomain())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.GoogleAuthResponse{
		AccessToken: tokenResp.AccessToken,
	})
}
