package redisrepository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/redis/go-redis/v9"
)

func (r *RedisRepository) EnqueueActivitiesUpsert(ctx context.Context, in domain.ActivityUpsertProcessRequest) error {
	// Marshal activities slice to JSON for Redis storage
	activitiesJSON, err := json.Marshal(in.Activities)
	if err != nil {
		errMsg := fmt.Sprintf("unable to marshal activities to JSON: %+v", err)
		return errors.New(errMsg)
	}

	id, err := r.redisCli.XAdd(ctx, &redis.XAddArgs{
		Stream: "activity-stream",
		Values: map[string]interface{}{
			"trip_id":    in.TripId,
			"trip_date":  in.Date.Format("2006-01-02"), // Format date as string
			"activities": string(activitiesJSON),       // Store as JSON string
		},
		MaxLen: 1000, // Max 1000 entries
		Approx: true,
	}).Result()
	log.Printf("enqueue id: %s", id)

	if err != nil {
		errMsg := fmt.Sprintf("unable to enqueue activities to redis: %+v", err)
		return errors.New(errMsg)
	}
	return nil
}
