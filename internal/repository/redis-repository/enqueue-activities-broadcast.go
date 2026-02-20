package redisrepository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *RedisRepository) EnqueueActivitiesBroadcast(ctx context.Context, tripId int, date time.Time) error {
	id, err := r.redisCli.XAdd(ctx, &redis.XAddArgs{
		Stream: "activity-broadcast-stream",
		Values: map[string]interface{}{
			"trip_id":   tripId,
			"trip_date": date.Format("2006-01-02"), // Format date as string
		},
		MaxLen: 1000, // Max 1000 entries
		Approx: true,
	}).Result()
	log.Printf("enqueue broadcast id: %s", id)

	if err != nil {
		errMsg := fmt.Sprintf("unable to enqueue activity broadcast to redis: %+v", err)
		return errors.New(errMsg)
	}
	return nil
}
