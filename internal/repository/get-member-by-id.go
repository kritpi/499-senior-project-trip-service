package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

func (r *Repository) GetMemberById(ctx context.Context, id string) (*domain.Member, error) {
	queryStr := fmt.Sprintf(
		`SELECT
			id,
			email,
			name,
			image_url
		FROM %s 
		WHERE id = @id`, r.cfg.Table.MemberTable)

	args := pgx.NamedArgs{
		"id": id,
	}

	var member entity.Member
	err := r.db.QueryRow(ctx, queryStr, args).Scan(
		&member.ID,
		&member.Email,
		&member.Name,
		&member.ImageUrl,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		log.Errorf("Repository.GetMemberByEmail: %v", err)
		return nil, err
	}

	return domain.Member{}.FromEntity(member), nil
}
