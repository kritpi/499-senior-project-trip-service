package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) GetTripExpenses(ctx context.Context, in domain.TripExpenseRequest) (*[]domain.ExpenseResponse, error) {
	queryString := fmt.Sprintf(`
		SELECT
			e.id AS expense_id,
			e.title,
			e.amount AS expense_amount,

			creator.name     AS created_by,       
			e.image_url,
			e.split_type,

			em.member_id,
			m.name           AS member_name,
			m.image_url      AS member_image,
			em.amount        AS member_amount,

			SUM(
				CASE
					WHEN em.member_id = @member_id
					THEN em.amount
					ELSE 0
				END
			) OVER (PARTITION BY e.id) AS my_shared

		FROM %s e
		JOIN %s em
			ON em.expense_id = e.id
		JOIN %s m
			ON m.id = em.member_id
		JOIN %s creator               
			ON creator.id = e.created_by

		WHERE e.trip_id = @trip_id
		ORDER BY e.id;
	`, r.cfg.Table.ExpenseTable, r.cfg.Table.ExpenseMemberTable, r.cfg.Table.MemberTable, r.cfg.Table.MemberTable)

	args := pgx.NamedArgs{
		"member_id": in.MemberId,
		"trip_id":   in.TripId,
	}

	expenseWithMember := make([]entity.ExpenseWithMemberRow, 0)

	rows, err := r.db.Query(ctx, queryString, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expenseRow entity.ExpenseWithMemberRow
		if err := rows.Scan(
			&expenseRow.ExpenseID,
			&expenseRow.Title,
			&expenseRow.ExpenseAmount,
			&expenseRow.CreatedBy,
			&expenseRow.ImageURL,
			&expenseRow.SplitType,
			&expenseRow.MemberID,
			&expenseRow.MemberName,
			&expenseRow.MemberImage,
			&expenseRow.MemberAmount,
			&expenseRow.MyShared,
		); err != nil {
			return nil, err
		}
		expenseWithMember = append(expenseWithMember, expenseRow)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	expenseResp := GroupExpense(expenseWithMember)

	return &expenseResp, nil
}

func GroupExpense(rows []entity.ExpenseWithMemberRow) []domain.ExpenseResponse {
	expenseMap := make(map[string]*domain.ExpenseResponse)

	for _, r := range rows {
		// Initialize expense if not exists
		if _, exists := expenseMap[r.ExpenseID]; !exists {
			expenseMap[r.ExpenseID] = &domain.ExpenseResponse{
				ExpenseId:   r.ExpenseID,
				Title:       r.Title,
				Amount:      r.ExpenseAmount,
				MyShared:    r.MyShared,
				CreatedBy:   r.CreatedBy,
				ImageUrl:    &r.ImageURL, // see note below
				SplitType:   r.SplitType,
				Participant: []domain.ExpenseMemberResponse{},
			}
		}

		// Append member
		expenseMap[r.ExpenseID].Participant = append(
			expenseMap[r.ExpenseID].Participant,
			domain.ExpenseMemberResponse{
				MemberId: r.MemberID,
				Name:     r.MemberName,
				ImageUrl: derefString(r.MemberImage),
				Amount:   r.MemberAmount,
			},
		)
	}

	expenses := make([]domain.ExpenseResponse, 0, len(expenseMap))
	for _, exp := range expenseMap {
		expenses = append(expenses, *exp)
	}

	return expenses
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
