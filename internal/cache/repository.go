package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	redis *redis.Client
}

func NewRepository(redis *redis.Client) *Repository {
	return &Repository{redis: redis}
}

func (r *Repository) Set(ctx context.Context, key string, value string, ttl int) error {
	var expiration time.Duration
	if ttl > 0 {
		expiration = time.Duration(ttl) * time.Second
	}
	return r.redis.Set(ctx, key, value, expiration).Err()
}

func (r *Repository) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

func (r *Repository) Delete(ctx context.Context, key string) error {
	return r.redis.Del(ctx, key).Err()
}
