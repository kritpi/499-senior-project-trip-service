package service

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) GetTripById(ctx context.Context, in domain.GetTripByIdRequest) (*domain.GetTripByIdResponse, error) {
	// Step 1: Check if member is part of the trip
	isMember, err := s.repo.CheckTripMembership(ctx, in.TripId, in.MemberId)
	if err != nil {
		log.Errorf("Service.GetTripById - CheckTripMembership: %v", err)
		return nil, err
	}

	if !isMember {
		log.Warnf("Service.GetTripById: member %s is not part of trip %d", in.MemberId, in.TripId)
		return nil, errors.New("forbidden: member not part of this trip")
	}

	// Step 2: Get trip details
	trip, err := s.repo.GetTripById(ctx, in.TripId)
	if err != nil {
		log.Errorf("Service.GetTripById - GetTripById: %v", err)
		return nil, err
	}

	if trip == nil {
		log.Warnf("Service.GetTripById: trip %d not found", in.TripId)
		return nil, errors.New("trip not found")
	}

	// Step 3: Get trip members
	members, err := s.repo.GetTripMembers(ctx, in.TripId)
	if err != nil {
		log.Errorf("Service.GetTripById - GetTripMembers: %v", err)
		return nil, err
	}

	return &domain.GetTripByIdResponse{
		Trip:    trip,
		Members: members,
	}, nil
}
