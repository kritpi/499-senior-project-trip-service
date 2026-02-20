package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/shopspring/decimal"
)

func (r *Repository) GetTripMemberTotalExpenses(ctx context.Context, in domain.TripExpenseRequest) (*decimal.Decimal, error) {
	queryString := fmt.Sprintf(`
		SELECT
			COALESCE(SUM(em.amount), 0) AS member_total_amount
		FROM
			%s em
		JOIN %s e ON e.id = em.expense_id
		WHERE e.trip_id = @trip_id 
		AND em.member_id = @member_id
	`, r.cfg.Table.ExpenseMemberTable, r.cfg.Table.ExpenseTable)

	args := pgx.NamedArgs{
		"trip_id":   in.TripId,
		"member_id": in.MemberId,
	}

	var memberTotalAmount decimal.Decimal
	err := r.db.QueryRow(ctx, queryString, args).Scan(&memberTotalAmount)
	if err != nil {
		return nil, err
	}
	return &memberTotalAmount, nil
}
