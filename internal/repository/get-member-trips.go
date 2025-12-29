package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) GetMemberTrips(ctx context.Context, in domain.GetMemberTripsRequest) (*domain.GetMemberTripsResponse, error) {
	queryString := fmt.Sprintf(`
		SELECT
			t.id,
			t.trip_name,
			t.start_date,
			t.end_date,
			t.main_location,
			tm.member_role
		FROM %s t
		INNER JOIN %s tm
  			ON tm.trip_id = t.id
		WHERE tm.member_id = @member_id
		ORDER BY t.start_date ASC;
	`, r.cfg.Table.TripTable, r.cfg.Table.TripMembersTable)

	args := pgx.NamedArgs{
		"member_id": in.MemberId,
	}

	var MemberTrips entity.GetMemberTripsResponse

	rows, err := r.db.Query(ctx, queryString, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trips := make([]entity.MemberTrips, 0)
	// trips := []entity.MemberTrips{}
	for rows.Next() {
		var t entity.MemberTrips
		if err := rows.Scan(
			&t.TripId,
			&t.TripName,
			&t.StartDate,
			&t.EndDate,
			&t.MainLocation,
			&t.Role,
		); err != nil {
			return nil, err
		}
		trips = append(trips, t)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	MemberTrips.Trips = trips

	return domain.GetMemberTripsResponse{}.FromEntity(MemberTrips), nil
}
