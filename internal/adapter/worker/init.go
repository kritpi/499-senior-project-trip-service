package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/kritpi/499-senior-project-trip-service/internal/adapter/worker/service"
	"github.com/redis/go-redis/v9"
)

type Worker struct {
	redis     *redis.Client
	upsert    *service.ActivityUpsertWorker
	broadcast *service.ActivityBroadcastWorker
}

func NewWorker(
	redis *redis.Client,
	upsert *service.ActivityUpsertWorker,
	broadcast *service.ActivityBroadcastWorker,
) *Worker {
	return &Worker{
		redis:     redis,
		upsert:    upsert,
		broadcast: broadcast,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	// Create error channel to receive errors from workers
	errChan := make(chan error, 2)

	// Start activity upsert worker
	go func() {
		log.Println("Starting Activity Upsert Worker...")
		err := w.upsert.ProcessStream(ctx)
		if err != nil {
			errChan <- fmt.Errorf("activity upsert worker error: %w", err)
		}
	}()

	// Start activity broadcast worker
	go func() {
		log.Println("Starting Activity Broadcast Worker...")
		err := w.broadcast.ProcessStream(ctx)
		if err != nil {
			errChan <- fmt.Errorf("activity broadcast worker error: %w", err)
		}
	}()

	// Wait for first error from either worker
	return <-errChan
}
