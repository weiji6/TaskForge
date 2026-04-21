package internal

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// CacheRepository is a struct that holds the Redis client.
type CacheRepository struct {
	client *redis.Client
}

// NewCacheRepository initializes a new CacheRepository.
func NewCacheRepository(redisAddr string) *CacheRepository {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &CacheRepository{client: client}
}

// Get retrieves a value from the cache.
func (r *CacheRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	return val, err
}

// Set sets a value in the cache.
func (r *CacheRepository) Set(ctx context.Context, key string, value string) error {
	return r.client.Set(ctx, key, value, 0).Err()
}