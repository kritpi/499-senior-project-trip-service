package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) GetActivities(ctx context.Context, in domain.GetActivitiesOfTheDayRequest) (*domain.GetActivitiesOfTheDayResponse, error) {
	queryString := fmt.Sprintf(`
		SELECT
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
		FROM %s WHERE
			trip_id = @trip_id
		AND
			activity_date = @activity_date
		ORDER BY rank ASC
	`, r.cfg.Table.ActivityTable)

	args := pgx.NamedArgs{
		"trip_id":       in.TripId,
		"activity_date": in.Date,
	}

	rows, err := r.db.Query(ctx, queryString, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []entity.Activity
	for rows.Next() {
		var ac entity.Activity
		err := rows.Scan(
			&ac.ID,
			&ac.TripId,
			&ac.ActivityDate,
			&ac.StartTime,
			&ac.EndTime,
			&ac.Note,
			&ac.Description,
			&ac.ActivityLocation, // pgx will automatically unmarshal JSON to the struct
			&ac.Category,
			&ac.Rank,
			&ac.CreatedAt,
			&ac.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		activities = append(activities, ac)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Convert entities to domain objects
	domainActivities := make([]domain.Activity, len(activities))
	for i, ac := range activities {
		domainActivities[i] = domain.Activity{
			ID:           ac.ID,
			TripId:       ac.TripId,
			ActivityDate: ac.ActivityDate,
			StartTime:    ac.StartTime,
			EndTime:      ac.EndTime,
			Note:         ac.Note,
			Description:  ac.Description,
			ActivityLocation: &domain.Location{
				Name:    ac.ActivityLocation.Name,
				Address: ac.ActivityLocation.Address,
				Lat:     ac.ActivityLocation.Lat,
				Lng:     ac.ActivityLocation.Lng,
			},
			Category:  ac.Category,
			Rank:      ac.Rank,
			CreatedAt: ac.CreatedAt,
			UpdatedAt: ac.UpdatedAt,
		}
	}

	return &domain.GetActivitiesOfTheDayResponse{
		Activities: domainActivities,
	}, nil
}
