package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/adapter/worker/model"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/port"
	"github.com/redis/go-redis/v9"
)

type ActivityUpsertWorker struct {
	svc   port.Service
	redis *redis.Client
}

func NewActivityUpsertWorker(
	svc port.Service,
	redis *redis.Client,
) *ActivityUpsertWorker {
	return &ActivityUpsertWorker{
		svc:   svc,
		redis: redis,
	}
}

func (w *ActivityUpsertWorker) ProcessStream(ctx context.Context) error {
	// Initialize the consumer group if it doesn't exist
	// Create the stream if it doesn't exist
	_, err := w.redis.XGroupCreateMkStream(ctx, "activity-stream", "activity-group", "0").Result()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		return err
	}

	for {
		msg, err := w.redis.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "activity-group",
			Consumer: "worker-1",
			Streams:  []string{"activity-stream", ">"},
			Count:    1,
			Block:    0,
		}).Result()
		log.Printf("%+v", msg)

		if err != nil {
			return err
		}

		for _, stream := range msg {
			for _, m := range stream.Messages {
				// Extract individual fields from Redis stream
				tripID, ok := m.Values["trip_id"].(string)
				if !ok {
					log.Printf("failed to get trip_id from message")
					continue
				}

				tripDate, ok := m.Values["trip_date"].(string)
				if !ok {
					log.Printf("failed to get trip_date from message")
					continue
				}

				activitiesJSON, ok := m.Values["activities"].(string)
				if !ok {
					log.Printf("failed to get activities from message")
					continue
				}

				// Unmarshal activities JSON
				var activities []domain.Activity
				err = json.Unmarshal([]byte(activitiesJSON), &activities)
				if err != nil {
					log.Printf("failed to unmarshal activities: %v", err)
					continue
				}

				// Parse trip_id as int
				var tripIDInt int
				_, err = fmt.Sscanf(tripID, "%d", &tripIDInt)
				if err != nil {
					log.Printf("failed to parse trip_id: %v", err)
					continue
				}

				// Parse trip_date
				parsedDate, err := time.Parse("2006-01-02", tripDate)
				if err != nil {
					log.Printf("failed to parse trip_date: %v", err)
					continue
				}

				// Construct the job
				job := model.ActivityUpsertProcessRequest{
					TripId:     tripIDInt,
					Date:       parsedDate,
					Activities: activities,
				}
				//call service
				_ = w.Handle(ctx, job)

				w.redis.XAck(ctx, "activity-stream", "activity-group", m.ID)
			}
		}
	}
}

func (w *ActivityUpsertWorker) Handle(
	ctx context.Context,
	job model.ActivityUpsertProcessRequest,
) error {

	req := domain.ActivityUpsertProcessRequest{
		TripId:     job.TripId,
		Date:       job.Date,
		Activities: job.Activities,
	}

	return w.svc.ActivityUpsertProcess(ctx, req)
}
