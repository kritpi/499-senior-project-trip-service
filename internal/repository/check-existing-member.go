package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) CheckExistingMember(ctx context.Context, checkExistingMember []string) (*[]domain.GetExistingMemberResp, error) {
	queryString := fmt.Sprintf(`
		SELECT
			id,
			email
		FROM %s
		WHERE email = ANY(@existingMember)
	`, r.cfg.Table.MemberTable)

	args := pgx.NamedArgs{
		"existingMember": checkExistingMember,
	}

	rows, err := r.db.Query(ctx, queryString, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var existing []entity.GetExistingMemberResp

	for rows.Next() {
		var id string
		var email string
		if err := rows.Scan(&id, &email); err != nil {
			return nil, err
		}
		existing = append(existing, entity.GetExistingMemberResp{
			ID:    id,
			Email: email,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convert entity to domain
	var result []domain.GetExistingMemberResp
	for _, e := range existing {
		result = append(result, *domain.GetExistingMemberResp{}.FromEntity(e))
	}

	return &result, nil
}
