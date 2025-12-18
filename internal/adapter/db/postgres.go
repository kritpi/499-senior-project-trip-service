package db

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context, connectionStr string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, connectionStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return pool
}
