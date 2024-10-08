package shortening

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	rdb *redis.Client
}

func NewRedisCache(rdb *redis.Client) *RedisCache {
	return &RedisCache{rdb: rdb}
}

func (r *RedisCache) Set(shortUrl, longUrl string, ttl time.Duration) error {
	return r.rdb.Set(context.Background(), shortUrl, longUrl, ttl).Err()
}

func (r *RedisCache) Get(shortUrl string) (string, error) {
	return r.rdb.Get(context.Background(), shortUrl).Result()
}

func (r *RedisCache) Delete(shortUrl string) error {
	return r.rdb.Del(context.Background(), shortUrl).Err()
}
