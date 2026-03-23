package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) InviteMember(ctx context.Context, in domain.TripInvitationRequest) (*domain.TripInvitationResponse, error) {
	now := time.Now().Local()

	// Check trip existing
	trip, err := s.repo.GetTripById(ctx, in.TripId)
	if err != nil || trip == nil {
		log.Errorf("unable to get trip id: %d error: %+v", in.TripId, err)
		return nil, err
	}
	if in.MemberId != trip.OwnerId {
		errMsg := fmt.Sprintf("unable to create invitation (unauthorized)")
		log.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}

	// Check invited member existing
	existedMember, err := s.repo.CheckExistingMember(ctx, in.Email)
	if err != nil {
		log.Errorf("unable to check member existing: %+v", err)
		return nil, err
	}
	if existedMember == nil {
		errMsg := fmt.Sprintf("member is not exists")
		return nil, errors.New(errMsg)
	}

	// Invite member
	err = s.repo.CreateTripMember(ctx, domain.CreateTripMemberRequest{
		TripId:    in.TripId,
		MemberId:  existedMember.ID,
		Role:      in.Role,
		CreatedAt: now,
	})
	if err != nil {
		log.Errorf("unable to invite member: %+v", err)
		return nil, err
	}
	returnedMember := domain.TripInvitationResponse {
		Email: existedMember.Email,
		Role: in.Role,
	}

	return &returnedMember, nil
}
