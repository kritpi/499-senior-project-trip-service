package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

func (r *Repository) GetTripExpenseTotalAmount(ctx context.Context, tripId int) (*decimal.Decimal, error) {
	queryString := fmt.Sprintf(`
		SELECT 
			COALESCE(SUM(amount), 0) AS total_amount
		FROM %s
		WHERE trip_id = @trip_id
	`, r.cfg.Table.ExpenseTable)

	args := pgx.NamedArgs{
		"trip_id": tripId,
	}

	var totalAmount decimal.Decimal
	err := r.db.QueryRow(ctx, queryString, args).Scan(&totalAmount)
	if err != nil {
		return nil, err
	}
	return &totalAmount, nil
}
