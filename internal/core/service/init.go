package service

import "github.com/kritpi/499-senior-project-trip-service/internal/core/port"

type service struct {
	repo port.Repository
}

func New(repo port.Repository) port.Service {
	return &service{
		repo: repo,
	}
}
