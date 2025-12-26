package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) UpsertTrip(ctx context.Context, in domain.UpsertTripRequest) (*domain.UpsertTripResponse, error) {
	args := pgx.NamedArgs{
		"id":            in.ID,
		"owner_id":      in.OwnerId,
		"trip_name":     in.TripName,
		"description":   in.Desciption,
		"start_date":    in.StartDate,
		"end_date":      in.EndDate,
		"main_location": in.MainLocation,
		"created_at":    in.CreatedAt,
		"updated_at":    in.UpdatedAt,
	}
	var queryString string

	if in.ID == nil {
		// INSERT
		queryString = fmt.Sprintf(`
            INSERT INTO %s (
                owner_id,
                trip_name,
                description,
                start_date,
                end_date,
                main_location,
                created_at,
                updated_at
            )
            VALUES (
                @owner_id,
                @trip_name,
                @description,
                @start_date,
                @end_date,
                @main_location,
                @created_at,
                @updated_at
            )
            RETURNING id;
        `, r.cfg.Table.TripTable)

	} else {
		// UPDATE
		queryString = fmt.Sprintf(`
            UPDATE %s SET 
                trip_name     = @trip_name,
                description   = @description,
                start_date    = @start_date,
                end_date      = @end_date,
                main_location = @main_location,
                updated_at    = @updated_at
            WHERE id = @id
            RETURNING id;
        `, r.cfg.Table.TripTable)
	}

	var id int
	if err := r.db.QueryRow(ctx, queryString, args).Scan(&id); err != nil {
		return nil, err
	}

	return &domain.UpsertTripResponse{TripId: id}, nil
}
