package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) GetTripById(ctx context.Context, tripId int) (*domain.Trip, error) {
	queryString := fmt.Sprintf(`
		SELECT
			id,
			owner_id,
			trip_name,
			description,
			start_date,
			end_date,
			main_location,
			image_url,
			created_at,
			updated_at
		FROM %s
		WHERE id = @id
	`, r.cfg.Table.TripTable)

	args := pgx.NamedArgs{
		"id": tripId,
	}

	var trip entity.Trip
	err := r.db.QueryRow(ctx, queryString, args).Scan(
		&trip.ID,
		&trip.OwnerId,
		&trip.TripName,
		&trip.Description,
		&trip.StartDate,
		&trip.EndDate,
		&trip.MainLocation,
		&trip.ImageUrl,
		&trip.CreatedAt,
		&trip.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		log.Errorf("Repository.GetTripById: %v", err)
		return nil, err
	}

	return domain.Trip{}.FromEntity(trip), nil
}
