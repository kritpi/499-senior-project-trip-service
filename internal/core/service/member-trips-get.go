package service

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) GetMemberTrips(ctx context.Context, in domain.GetMemberTripsRequest) (*domain.GetMemberTripsResponse, error) {
	trips, err := s.repo.GetMemberTrips(ctx, in)
	if err != nil {
		log.Errorf("unable to get trips from this member: %s, error: %+v", in.MemberId, err)
		return nil, err
	}
	if trips == nil {
		return nil, nil
	}
	return trips, nil
}
