package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) DeleteInvitedMember(ctx context.Context, memberId string, tripId int) error {
	queryString := fmt.Sprintf(`
		DELETE FROM %s
		WHERE 
			member_id = @member_id 
		AND
			trip_id = @trip_id
	`, r.cfg.Table.TripMembersTable)

	args := pgx.NamedArgs{
		"member_id": memberId,
		"trip_id":   tripId,
	}

	cmdTag, err := r.db.Exec(ctx, queryString, args)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		errMsg := fmt.Sprintf("rows not found")
		return errors.New(errMsg)
	}
	return nil
}
