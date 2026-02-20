package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) DeleteExpenseMember(ctx context.Context, expenseId string) error {
	queryString := fmt.Sprintf(`
		DELETE FROM %s WHERE expense_id = @expense_id
	`, r.cfg.Table.ExpenseMemberTable)

	args := pgx.NamedArgs{
		"expense_id": expenseId,
	}

	_, err := r.db.Exec(ctx, queryString, args)
	if err != nil {
		return err
	}
	return nil
}
