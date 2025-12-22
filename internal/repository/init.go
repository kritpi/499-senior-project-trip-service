package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kritpi/499-senior-project-trip-service/property"
)

type Repository struct {
	db  *pgxpool.Pool
	cfg property.Property
}

func New(ctx context.Context, db *pgxpool.Pool, cfg property.Property) Repository {
	return Repository{
		db:  db,
		cfg: cfg,
	}
}
