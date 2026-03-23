package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) ActivityUpsertBroadcast(ctx context.Context, tripId int, date time.Time) (*domain.ActivityJoinResponse, error) {
	// Query activities by trip ID and date
	activities, err := s.repo.GetActivities(ctx, domain.GetActivitiesOfTheDayRequest{
		TripId: tripId,
		Date:   date,
	})
	if err != nil {
		errMsg := fmt.Sprintf("unable to get activities for broadcast: %+v", err)
		return nil, errors.New(errMsg)
	}

	// Return the activities in the expected response format
	dateStr := date.Format("2006-01-02")
	return &domain.ActivityJoinResponse{
		TripId:     tripId,
		Date:       dateStr,
		Activities: activities.Activities,
		IsEditable: false, // Broadcast is read-only
	}, nil
}
