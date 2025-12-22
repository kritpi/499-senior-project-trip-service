package repository

import (
	"context"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) GetTripsInfo(ctx context.Context) ([]domain.Trip, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, destination, start_date, end_date, status, created_at, updated_at FROM trips")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []domain.Trip
	for rows.Next() {
		var trip domain.Trip
		if err := rows.Scan(&trip.ID, &trip.Name, &trip.Destination, &trip.StartDate, &trip.EndDate, &trip.Status, &trip.CreatedAt, &trip.UpdatedAt); err != nil {
			return nil, err
		}
		trips = append(trips, trip)
	}

	return trips, rows.Err()
}
