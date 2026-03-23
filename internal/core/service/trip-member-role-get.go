package service

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) GetTripMemberRole(ctx context.Context, in domain.GetTripMemberRoleRequest) (*domain.GetTripMemberRoleResponse, error) {
	_, err := s.repo.GetTripMemberRole(ctx, in.MemberId, in.TripId)
	if err != nil {
		log.Errorf("unable to get trip member role: %+v", err)
		return nil, err
	}
	return nil, nil
	
}
