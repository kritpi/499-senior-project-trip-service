package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) CheckExistingMember(ctx context.Context, email string) (*domain.GetExistingMemberResp, error) {
	queryString := fmt.Sprintf(`
		SELECT
			id,
			email
		FROM %s
		WHERE email = @email
	`, r.cfg.Table.MemberTable)

	args := pgx.NamedArgs{
		"email": email,
	}

	var member entity.GetExistingMemberResp
	err := r.db.QueryRow(ctx, queryString, args).Scan(&member.ID, &member.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Member not found
		}
		return nil, err
	}

	return domain.GetExistingMemberResp{}.FromEntity(member), nil
}
