package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/port"
	"github.com/kritpi/499-senior-project-trip-service/property"
)

type RestHandler interface {
	GetTrips(c *fiber.Ctx) error

	// Authentication
	GoogleAuth(c *fiber.Ctx) error
}

type restHandler struct {
	svc port.Service
	cfg property.Property
}

func NewRestApi(svc port.Service, cfg property.Property) RestHandler {
	return &restHandler{
		svc: svc,
		cfg: cfg,

	}
}
