package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

func (r *Repository) GetTripMemberRole(ctx context.Context, memberId string, tripId int) (*enum.MemberRole, error) {
	queryString := fmt.Sprintf(`
		SELECT 
			member_role 
		FROM %s 
		WHERE 
			member_id = @member_id 
		AND 
			trip_id = @trip_id
		`, r.cfg.Table.TripMembersTable)

	args := pgx.NamedArgs{
		"member_id": memberId,
		"trip_id":   tripId,
	}

	var memberRole enum.MemberRole
	err := r.db.QueryRow(ctx, queryString, args).Scan(&memberRole)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		log.Errorf("Repository.GetTripMemberRole: %+v", err)
		return nil, err
	}
	return &memberRole, nil
}
