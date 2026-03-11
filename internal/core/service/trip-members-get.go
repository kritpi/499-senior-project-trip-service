package service

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) GetTripMembers(ctx context.Context, in domain.GetTripMembersRequest) ([]domain.TripMemberDetails, error) {
	// Check if member is part of the trip
	isMember, err := s.repo.CheckTripMembership(ctx, in.TripId, in.MemberId)
	if err != nil {
		log.Errorf("Service.GetTripMembers - CheckTripMembership: %v", err)
		return nil, err
	}

	if !isMember {
		log.Warnf("Service.GetTripMembers: member %s is not part of trip %d", in.MemberId, in.TripId)
		return nil, errors.New("forbidden: member not part of this trip")
	}

	members, err := s.repo.GetTripMembers(ctx, in.TripId)
	if err != nil {
		log.Errorf("Service.GetTripMembers - GetTripMembers: %v", err)
		return nil, err
	}

	return members, nil
}
