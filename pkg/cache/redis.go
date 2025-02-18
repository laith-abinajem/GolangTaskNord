package cache

import (
	"context"
	"task/pkg/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Url,
		// MinIdleConns: cfg.Redis.MinIdleConn,
		// PoolSize:     cfg.Redis.PoolSize,
		// PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		//Password:     cfg.Redis.RedisPassword, // no password set
		//DB:           cfg.Redis.DB,            // use default DB
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
