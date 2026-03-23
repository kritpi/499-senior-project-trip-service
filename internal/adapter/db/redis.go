package db

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/property"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context, cfg property.Property) *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host + ":" + cfg.Redis.Port,
	})

	// Test the connection
	if err := redisCli.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	return redisCli
}
