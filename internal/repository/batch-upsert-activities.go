package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) BatchUpsertActivities(ctx context.Context, in []domain.Activity) error {
	queryString := fmt.Sprintf(`
		INSERT INTO %s (
			id,
			trip_id,
			activity_date,
			start_time,
			end_time,
			note,
			description,
			activity_location,
			category,
			rank,
			created_at,
			updated_at
		)
		VALUES (
			@id,
			@trip_id,
			@activity_date,
			@start_time,
			@end_time,
			@note,
			@description,
			@activity_location::jsonb,
			@category,
			@rank,
			@created_at,
			@updated_at
		) ON CONFLICT (id) DO UPDATE SET
			trip_id = EXCLUDED.trip_id,
			activity_date = EXCLUDED.activity_date,
			start_time = EXCLUDED.start_time,
			end_time = EXCLUDED.end_time,
			note = EXCLUDED.note,
			description = EXCLUDED.description,
			activity_location = EXCLUDED.activity_location,
			category = EXCLUDED.category,
			rank = EXCLUDED.rank,
			updated_at = EXCLUDED.updated_at
	`, r.cfg.Table.ActivityTable)

	// Use batch for efficient bulk operations
	batch := &pgx.Batch{}

	for _, activity := range in {
		// Convert domain.Location to JSON string for ::jsonb cast
		var activityLocationStr interface{}
		if activity.ActivityLocation != nil {
			locationEntity := entity.Location{
				Name:    activity.ActivityLocation.Name,
				Address: activity.ActivityLocation.Address,
				Lat:     activity.ActivityLocation.Lat,
				Lng:     activity.ActivityLocation.Lng,
			}
			locationJSON, err := json.Marshal(locationEntity)
			if err != nil {
				return fmt.Errorf("failed to marshal activity_location for activity %s: %w", activity.ID, err)
			}
			activityLocationStr = string(locationJSON)
		}

		args := pgx.NamedArgs{
			"id":                activity.ID,
			"trip_id":           activity.TripId,
			"activity_date":     activity.ActivityDate,
			"start_time":        activity.StartTime,
			"end_time":          activity.EndTime,
			"note":              activity.Note,
			"description":       activity.Description,
			"activity_location": activityLocationStr, // Will be cast to jsonb in SQL
			"category":          activity.Category,
			"rank":              activity.Rank,
			"created_at":        activity.CreatedAt,
			"updated_at":        activity.UpdatedAt,
		}

		batch.Queue(queryString, args)
	}

	// Execute the batch
	batchResults := r.db.SendBatch(ctx, batch)
	defer batchResults.Close()

	// Check each result for errors
	for i := 0; i < len(in); i++ {
		_, err := batchResults.Exec()
		if err != nil {
			return fmt.Errorf("failed to upsert activity at index %d (id: %s): %w", i, in[i].ID, err)
		}
	}

	return nil
}
