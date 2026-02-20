package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) BatchInsertExpenseMember(ctx context.Context, in domain.ExpenseMemberRequest) error {
	queryString := fmt.Sprintf(`
		INSERT INTO %s (
			expense_id,
			member_id,
			amount
		) VALUES (
			@expense_id,
			@member_id,
			@amount 
		)
	`, r.cfg.Table.ExpenseMemberTable)

	batch := &pgx.Batch{}

	for _, member := range in.ExpenseMember {
		args := pgx.NamedArgs{
			"expense_id": in.ExpenseId,
			"member_id":  member.MemberId,
			"amount":     member.Amount,
		}

		batch.Queue(queryString, args)
	}

	batchResults := r.db.SendBatch(ctx, batch)
	defer batchResults.Close()

	for i := 0; i < len(in.ExpenseMember); i++ {
		_, err := batchResults.Exec()
		if err != nil {
			return fmt.Errorf("failed to batch insert expense members: %+v", err)
		}
	}

	return nil
}
