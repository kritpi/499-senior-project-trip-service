package redisrepository

import (
	"context"

	"github.com/kritpi/499-senior-project-trip-service/property"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redisCli *redis.Client
	cfg      property.Property
}

func New(ctx context.Context, redisCli *redis.Client, cfg property.Property) RedisRepository {
	return RedisRepository{
		redisCli: redisCli,
		cfg:      cfg,
	}
}
