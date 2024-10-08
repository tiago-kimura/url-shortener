package shortening

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache interface {
	Get(urlId string) (string, error)
	Set(urlId, url_original string, ttl time.Duration) error
	Delete(urlId string) error
}

type RedisCacheImpl struct {
	Rdb *redis.Client
}

func NewRedisCache(rdb *redis.Client) *RedisCacheImpl {
	return &RedisCacheImpl{Rdb: rdb}
}

func (r *RedisCacheImpl) Set(urlId, url_original string, ttl time.Duration) error {
	return r.Rdb.Set(context.Background(), urlId, url_original, ttl).Err()
}

func (r *RedisCacheImpl) Get(urlId string) (string, error) {
	return r.Rdb.Get(context.Background(), urlId).Result()
}

func (r *RedisCacheImpl) Delete(urlId string) error {
	return r.Rdb.Del(context.Background(), urlId).Err()
}
