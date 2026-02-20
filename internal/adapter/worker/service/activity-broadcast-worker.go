package service

import (
	"context"
	"fmt"
	"log"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/port"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/socket"
	"github.com/redis/go-redis/v9"
)

type ActivityBroadcastWorker struct {
	svc          port.Service
	redis        *redis.Client
	socketServer *socketio.Server
}

func NewActivityBroadcastWorker(
	svc port.Service,
	redis *redis.Client,
	socketServer *socketio.Server,
) *ActivityBroadcastWorker {
	return &ActivityBroadcastWorker{
		svc:          svc,
		redis:        redis,
		socketServer: socketServer,
	}
}

func (w *ActivityBroadcastWorker) ProcessStream(ctx context.Context) error {
	// Initialize the consumer group if it doesn't exist
	_, err := w.redis.XGroupCreateMkStream(ctx, "activity-broadcast-stream", "activity-broadcast-group", "0").Result()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		return err
	}

	for {
		msg, err := w.redis.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "activity-broadcast-group",
			Consumer: "worker-1",
			Streams:  []string{"activity-broadcast-stream", ">"},
			Count:    1,
			Block:    0,
		}).Result()
		log.Printf("Broadcast message received: %+v", msg)

		if err != nil {
			return err
		}

		for _, stream := range msg {
			for _, m := range stream.Messages {
				// Extract trip_id and trip_date from Redis stream
				tripIDStr, ok := m.Values["trip_id"].(string)
				if !ok {
					log.Printf("failed to get trip_id from message")
					continue
				}

				tripDate, ok := m.Values["trip_date"].(string)
				if !ok {
					log.Printf("failed to get trip_date from message")
					continue
				}

				// Parse trip_id as int
				var tripID int
				_, err = fmt.Sscanf(tripIDStr, "%d", &tripID)
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

				// Call service to get activities
				activitiesResp, err := w.svc.ActivityUpsertBroadcast(ctx, tripID, parsedDate)
				if err != nil {
					log.Printf("failed to get activities for broadcast: %v", err)
					w.redis.XAck(ctx, "activity-broadcast-stream", "activity-broadcast-group", m.ID)
					continue
				}

				// Broadcast to socket room
				roomName := socket.GetActivityRoomName(tripID, tripDate)
				payload := dto.ActivityResponse{}.FromDomain(activitiesResp)

				w.socketServer.BroadcastToRoom("/", roomName, socket.EventActivityBroadcast, payload)
				log.Printf("Broadcasted %d activities to room: %s", len(activitiesResp.Activities), roomName)

				// Acknowledge the message
				w.redis.XAck(ctx, "activity-broadcast-stream", "activity-broadcast-group", m.ID)
			}
		}
	}
}
