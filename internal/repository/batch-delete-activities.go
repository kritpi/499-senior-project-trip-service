package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) BatchDeleteActivities(ctx context.Context, in []domain.Activity) error {
	if len(in) == 0 {
		return nil // Nothing to delete
	}

	queryString := fmt.Sprintf(`
		DELETE FROM %s WHERE id = @id
	`, r.cfg.Table.ActivityTable)

	batch := &pgx.Batch{}

	for _, activity := range in {
		args := pgx.NamedArgs{
			"id": activity.ID,
		}
		batch.Queue(queryString, args)
	}

	batchResults := r.db.SendBatch(ctx, batch)
	defer batchResults.Close()

	for i := 0; i < len(in); i++ {
		_, err := batchResults.Exec()
		if err != nil {
			return fmt.Errorf("failed to delete activity at index %d (id: %s): %w", i, in[i].ID, err)
		}
	}
	return nil
}
