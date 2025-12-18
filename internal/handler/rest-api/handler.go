package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/port"
)

type RestHandler interface {
	GetTrips(c *fiber.Ctx) error
}

type restHandler struct {
	svc port.Service
}

func NewRestApi(svc port.Service) RestHandler {
	return &restHandler{
		svc: svc,
	}
}
