package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) GetTripMembers(ctx context.Context, tripId int) ([]domain.TripMemberDetails, error) {
	queryString := fmt.Sprintf(`
		SELECT
			m.id as member_id,
			m.email,
			m.name,
			m.image_url,
			tm.member_role
		FROM %s tm
		INNER JOIN %s m ON tm.member_id = m.id
		WHERE tm.trip_id = @trip_id
		ORDER BY tm.member_role, m.name
	`, r.cfg.Table.TripMembersTable, r.cfg.Table.MemberTable)

	args := pgx.NamedArgs{
		"trip_id": tripId,
	}

	rows, err := r.db.Query(ctx, queryString, args)
	if err != nil {
		log.Errorf("Repository.GetTripMembers: %v", err)
		return nil, err
	}
	defer rows.Close()

	var members []domain.TripMemberDetails
	for rows.Next() {
		var member entity.TripMemberWithDetails
		err := rows.Scan(
			&member.MemberId,
			&member.Email,
			&member.Name,
			&member.ImageUrl,
			&member.Role,
		)
		if err != nil {
			log.Errorf("Repository.GetTripMembers - Scan: %v", err)
			return nil, err
		}
		members = append(members, domain.TripMemberDetails{}.FromEntity(member))
	}

	if err := rows.Err(); err != nil {
		log.Errorf("Repository.GetTripMembers - Rows: %v", err)
		return nil, err
	}

	return members, nil
}
