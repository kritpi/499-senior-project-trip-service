package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) ActivityUpsertProcess(ctx context.Context, in domain.ActivityUpsertProcessRequest) error {
	oldActivities, err := s.repo.GetActivities(ctx, domain.GetActivitiesOfTheDayRequest{
		TripId: in.TripId,
		Date:   in.Date,
	})
	if err != nil {
		errMsg := fmt.Sprintf("unable to get previous activities: %+v", err)
		return errors.New(errMsg)
	}
	// make old activities map
	oldActivitiesMap := make(map[string]domain.Activity)
	for _, act := range oldActivities.Activities {
		oldActivitiesMap[act.ID] = act
	}

	//make new activities map
	newActivitiesMap := make(map[string]domain.Activity)
	for _, act := range in.Activities {
		if act.ID == "" {
			act.ID = uuid.New().String()
		}
		newActivitiesMap[act.ID] = act
	}

	// get activities diff
	activitiesToUpsert, activitiesToDelete := diffActivities(&oldActivitiesMap, &newActivitiesMap)

	txErr := s.repo.Transactional(ctx, func(txCtx context.Context) error {
		err := s.repo.BatchUpsertActivities(txCtx, *activitiesToUpsert)
		if err != nil {
			return err
		}
		err = s.repo.BatchDeleteActivities(txCtx, *activitiesToDelete)
		if err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		errMsg := fmt.Sprintf("transactional error: %+v", txErr)
		return errors.New(errMsg)
	}
	err = s.redisRepo.EnqueueActivitiesBroadcast(ctx, in.TripId, in.Date)
	if err != nil {
		errMsg := fmt.Sprintf("failed to enqueue broadcast: %+v", err)
		return errors.New(errMsg)
	}

	return nil
}

func diffActivities(oldActivities, newActivities *map[string]domain.Activity) (toUpsert, toDelete *[]domain.Activity) {
	toUpsertSlice := []domain.Activity{}
	toDeleteSlice := []domain.Activity{}
	now := time.Now()

	// All new activities should be upserted (whether new or existing IDs)
	for _, newAct := range *newActivities {
		if oldAct, exists := (*oldActivities)[newAct.ID]; exists {
			// Existing activity: preserve CreatedAt, update UpdatedAt
			newAct.CreatedAt = oldAct.CreatedAt
			newAct.UpdatedAt = now
		} else {
			// New activity: set both timestamps to current time
			newAct.CreatedAt = now
			newAct.UpdatedAt = now
		}
		toUpsertSlice = append(toUpsertSlice, newAct)
	}

	// Any old activity not in new activities should be deleted
	for _, oldAct := range *oldActivities {
		if _, exists := (*newActivities)[oldAct.ID]; !exists {
			toDeleteSlice = append(toDeleteSlice, oldAct)
		}
	}

	return &toUpsertSlice, &toDeleteSlice
}
