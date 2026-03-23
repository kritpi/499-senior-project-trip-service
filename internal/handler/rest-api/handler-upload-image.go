package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

func (h *restHandler) UploadImage(c *fiber.Ctx) error {
	ctx := c.Context()
	jwtClaims := c.Locals("user").(*model.JWTCustomClaims)

	// Get file from form
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "image file is required",
		})
	}

	// Validate file type (accept only images)
	contentType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "only image files are allowed",
		})
	}

	// Call service
	req := domain.UploadImageRequest{
		File:        file,
		ContentType: contentType,
		MemberId:    jwtClaims.ID,
	}

	resp, err := h.svc.UploadImage(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Convert to DTO and return
	return c.Status(fiber.StatusOK).JSON(dto.UploadImageResponse{}.FromDomain(resp))
}
