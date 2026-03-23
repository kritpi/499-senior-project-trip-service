package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) DeleteExpense(ctx context.Context, expenseId string) error {
	queryString := fmt.Sprintf(`
		DELETE FROM %s WHERE id = @id
	`, r.cfg.Table.ExpenseTable)

	args := pgx.NamedArgs{
		"id": expenseId,
	}

	_, err := r.db.Exec(ctx, queryString, args)
	if err != nil {
		return err
	}
	return nil
}
