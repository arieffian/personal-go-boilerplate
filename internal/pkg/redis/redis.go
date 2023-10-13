package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService interface {
	SetCacheWithExpiration(ctx context.Context, key string, value interface{}, duration int) error
	SetCacheWithoutExpiration(ctx context.Context, key string, value interface{}) error
	GetCache(ctx context.Context, key string, result interface{}) error
}

type RedisConfig struct {
	Host string
	Port int
	TTL  int
}

type RedisClient struct {
	ttl    int
	client *redis.Client
}

var _ RedisService = (*RedisClient)(nil)

func (r *RedisClient) SetCacheWithExpiration(ctx context.Context, key string, value interface{}, duration int) error {
	ttl := duration
	if ttl == 0 {
		ttl = r.ttl
	}

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, key, jsonValue, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) SetCacheWithoutExpiration(ctx context.Context, key string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, key, jsonValue, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) GetCache(ctx context.Context, key string, result interface{}) error {
	val, err := r.client.Get(ctx, "key").Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		return err
	}

	return nil
}

func NewRedisConnection(cfg RedisConfig) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "",
		DB:       0,
	})

	return &RedisClient{
		ttl:    cfg.TTL,
		client: rdb,
	}
}
