package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type redisConfig interface {
	RedisHost() string
	RedisPort() string
	RedisDB() int
	CacheTimeout() time.Duration
}

func NewRedis(cfg redisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%s", cfg.RedisHost(), cfg.RedisPort()),
		DB:          cfg.RedisDB(),
		DialTimeout: cfg.CacheTimeout(),
		ReadTimeout: cfg.CacheTimeout(),
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
