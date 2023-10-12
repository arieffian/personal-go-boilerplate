package database

import (
	"context"
	"fmt"

	_ "github.com/lib/pq" //need to running query: postgres driver
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host string
	Port int
	TTL  int
}

type redisConfig struct {
	redisConfig RedisConfig
}

func (d *redisConfig) CreateRedisClient(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", d.redisConfig.Host, d.redisConfig.Port),
		Password: "",
		DB:       0,
	})

	return rdb
}

func NewRedisConnection(cfg RedisConfig) *redisConfig {
	return &redisConfig{redisConfig: cfg}
}
