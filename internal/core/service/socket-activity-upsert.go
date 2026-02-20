package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

func (s *service) SocketActivityUpsert(ctx context.Context, in domain.ActivityUpsertRequest) error {
	// Verify member has permission to edit
	MemberRole, err := s.repo.GetTripMemberRole(ctx, in.MemberId, in.TripId)
	if err != nil {
		errMsg := fmt.Sprintf("unable to get member role: %+v", err)
		return errors.New(errMsg)
	}

	if MemberRole == nil {
		return errors.New("member role is nil - user may not have access to this trip")
	}

	// Only editors and owners can upsert activities
	if *MemberRole == enum.MemberRoleViewer {
		return errors.New("insufficient permissions - viewers cannot edit activities")
	}

	// enqueue activities request to redis stream
	err = s.redisRepo.EnqueueActivitiesUpsert(ctx, domain.ActivityUpsertProcessRequest{
		TripId:     in.TripId,
		Date:       in.Date,
		Activities: in.Activities,
	})
	if err != nil {
		errMsg := fmt.Sprintf("unable to enqueue activities to redis: %+v", err)
		return errors.New(errMsg)
	}

	return nil
}
