package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) CreateMember(ctx context.Context, member domain.Member) (id string, err error) {
	queryStr := fmt.Sprintf(
		`INSERT INTO %s (
			id,
			email,
			name,
			image_url
		) VALUES (
		 @id,
		 @email,
		 @name,
		 @image_url
		) RETURNING id`, r.cfg.Table.MemberTable)

	args := pgx.NamedArgs{
		"id":        member.ID,
		"email":     member.Email,
		"name":      member.Name,
		"image_url": member.ImageUrl,
	}

	var insertedID string
	if err := r.db.QueryRow(ctx, queryStr, args).Scan(&insertedID); err != nil {
		log.Errorf("create member error: %v", err)
		return "", err
	}

	log.Infof("Inserted member id=%s", insertedID)
	return insertedID, nil
}
