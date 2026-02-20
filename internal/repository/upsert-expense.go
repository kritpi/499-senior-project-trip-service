package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) UpsertExpense(ctx context.Context, in domain.UpsertExpenseRequest) error {
	queryString := fmt.Sprintf(`
		INSERT INTO %s (
			id,
			trip_id,
			title,
			amount,
			created_by,
			split_type,
			image_url
		) 
		VALUES (
			@id,
			@trip_id,
			@title,
			@amount,
			@created_by,
			@split_type,
			@image_url
		) ON CONFLICT (id) DO UPDATE SET
		 	title = EXCLUDED.title,
		 	amount = EXCLUDED.amount,
		 	created_by = EXCLUDED.created_by,
		 	split_type = EXCLUDED.split_type,
		 	image_url = EXCLUDED.image_url
	`, r.cfg.Table.ExpenseTable)

	args := pgx.NamedArgs{
		"id":         in.ExpenseId,
		"trip_id":    in.TripId,
		"title":      in.Title,
		"amount":     in.Amount,
		"created_by": in.CreatedBy,
		"split_type": in.SplitType,
		"image_url":  in.ImageUrl,
	}

	_, err := r.db.Exec(ctx, queryString, args)
	if err != nil {
		return err
	}
	return nil
}
