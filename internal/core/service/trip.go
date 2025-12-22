package service

import (
	"context"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) GetTrips(ctx context.Context) ([]domain.Trip, error) {
	trips, err := s.repo.GetTripsInfoTest(ctx)
	if err != nil {
		return nil, err
	}
	return trips, nil
}
