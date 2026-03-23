package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

func (s *service) SocketActivityJoin(ctx context.Context, in domain.ActivityJoinRequest) (*domain.ActivityJoinResponse, error) {
	MemberRole, err := s.repo.GetTripMemberRole(ctx, in.MemberId, in.TripId)
	if err != nil {
		errMsg := fmt.Sprintf("unable to get member role: %+v", err)
		return nil, errors.New(errMsg)
	}

	// Check for nil pointer to prevent panic
	if MemberRole == nil {
		return nil, errors.New("member role is nil - user may not have access to this trip")
	}

	// Query activities of the day from database
	activitiesResp, err := s.repo.GetActivities(ctx, domain.GetActivitiesOfTheDayRequest{
		TripId: in.TripId,
		Date:   in.TripDate,
	})
	if err != nil {
		errMsg := fmt.Sprintf("unable to get activities: %+v", err)
		return nil, errors.New(errMsg)
	}

	// Format date as string (YYYY-MM-DD)
	dateStr := in.TripDate.Format("2006-01-02")

	resp := domain.ActivityJoinResponse{
		TripId:     in.TripId,
		Date:       dateStr,
		Activities: activitiesResp.Activities,
		IsEditable: false,
	}

	switch *MemberRole {
	case enum.MemberRoleViewer:
		resp.IsEditable = false
	default:
		resp.IsEditable = true
	}

	return &resp, nil
}
