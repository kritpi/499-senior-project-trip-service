package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) CreateTripMember(ctx context.Context, in domain.CreateTripMemberRequest) error {
	queryString := fmt.Sprintf(`
        INSERT INTO %s (
            trip_id,
            member_id,
            member_role,
            created_at
        ) VALUES (
            @trip_id,
            @member_id,
            @member_role,
            @created_at 
        )
    `, r.cfg.Table.TripMembersTable)

	args := pgx.NamedArgs{
		"trip_id":     in.TripId,
		"member_id":   in.MemberId,
		"member_role": in.Role,
		"created_at":  in.CreatedAt,
	}

	_, err := r.db.Exec(ctx, queryString, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return errors.New("trip member already exists")
			}
		}
		return err
	}
	return nil
}
