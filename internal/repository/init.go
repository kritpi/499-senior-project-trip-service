package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, db *pgxpool.Pool) Repository {
	return Repository{
		db: db,
	}
}
