package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) BatchCreateTripMember(ctx context.Context, in domain.BatchCreateTripMemberRequest) error {
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

    tx, err := r.db.Begin(ctx)
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)

    batch := &pgx.Batch{}

    for _, m := range in.Members {
        batch.Queue(queryString, pgx.NamedArgs{
            "trip_id":     in.TripId,
            "member_id":   m.MemberId,
            "member_role": m.Role,
            "created_at":  in.CreatedAt,
        })
    }

    br := tx.SendBatch(ctx, batch)

    for range in.Members {
        if _, err := br.Exec(); err != nil {
            br.Close()
            return err
        }
    }

    if err := br.Close(); err != nil {
        return err
    }

    return tx.Commit(ctx)
}
