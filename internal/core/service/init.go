package service

import (
	"github.com/kritpi/499-senior-project-trip-service/internal/core/port"
	"github.com/kritpi/499-senior-project-trip-service/property"
)

type service struct {
	repo port.Repository
	cfg  property.Property
}

func New(repo port.Repository, cfg *property.Property) port.Service {
	return &service{
		repo: repo,
		cfg:  *cfg,
	}
}
