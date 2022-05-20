package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis(cfg *RedisConfig, logger *zap.Logger) (*redis.Client, error) {
	if cfg == nil {
		return nil, nil
	}
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass, // no password set
		DB:       cfg.DB,   // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("redis connect ping failed, err:", zap.Error(err))
		return nil, err
	}
	logger.Info("redis connect ping response:", zap.String("pong", pong))
	return client, nil
}
