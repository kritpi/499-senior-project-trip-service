package service

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

func (s *service) InviteMember(ctx context.Context, in domain.TripInvitationRequest) (*domain.InviteMemberResponse, error) {
	now := time.Now().Local()
	// validate if trip is valid
	trip, err := s.repo.GetTripById(ctx, in.TripId)
	if err != nil {
		log.Errorf("unable to get existing trip: %+v", err)
		return nil, err
	}
	if trip == nil {
		errMsg := "trip is not exist"
		log.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}
	// validate if trip owner valid
	if trip.OwnerId != in.MemberId {
		errMsg := "only the trip owner is authorized to invite members"
		log.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}
	// validate if invited member valid
	invitedMemberMap := make(map[string]string)
	checkExistingMember := make([]string, len(in.Member))

	for i, m := range in.Member {
		invitedMemberMap[m.Email] = string(m.Role)
		checkExistingMember[i] = m.Email
	}

	validatedMember, err := s.repo.CheckExistingMember(ctx, checkExistingMember)
	if err != nil {
		log.Errorf("unable to check existing member error: %+v", err)
		return nil, err
	}
	
	createTripMember := make([]domain.TripMember, 0)
	for _, m := range *validatedMember {
		// remove member from map
		createTripMember = append(createTripMember, domain.TripMember{
			MemberId: m.ID,
			Role:     enum.MemberRole(invitedMemberMap[m.Email]),
		})
		delete(invitedMemberMap, m.Email)
	}

	// write invited member to db
	err = s.repo.BatchCreateTripMember(ctx, domain.BatchCreateTripMemberRequest{
		TripId:    in.TripId,
		Members:   createTripMember,
		CreatedAt: now,
	})

	uninvitedMember := make([]string, 0)
	for key := range invitedMemberMap {
		uninvitedMember = append(uninvitedMember, key)
	}

	// return uninvited member
	resp := domain.InviteMemberResponse{
		Message: "members not existing",
		Member:  uninvitedMember,
	}

	return &resp, nil
}
