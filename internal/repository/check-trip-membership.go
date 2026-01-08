package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) CheckTripMembership(ctx context.Context, tripId int, memberId string) (bool, error) {
	queryString := fmt.Sprintf(`
		SELECT EXISTS(
			SELECT 1
			FROM %s
			WHERE trip_id = @trip_id AND member_id = @member_id
		)
	`, r.cfg.Table.TripMembersTable)

	args := pgx.NamedArgs{
		"trip_id":   tripId,
		"member_id": memberId,
	}

	var exists bool
	err := r.db.QueryRow(ctx, queryString, args).Scan(&exists)
	if err != nil {
		log.Errorf("Repository.CheckTripMembership: %v", err)
		return false, err
	}

	return exists, nil
}
