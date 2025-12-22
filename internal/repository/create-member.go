package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) CreateMember(ctx context.Context, member domain.Member) error {
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
		)`, r.cfg.Table.MemberTable)

	args := pgx.NamedArgs{
		"id":        member.ID,
		"email":     member.Email,
		"name":      member.Name,
		"image_url": member.ImageUrl,
	}

	c, err := r.db.Exec(ctx, queryStr, args)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	log.Infof("Inserted: %+v", c.RowsAffected())
	return nil
}
