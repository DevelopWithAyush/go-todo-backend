package cache

import (
	"context"
	"github.com/developwithayush/go-todo-app/internal/config"
	"github.com/developwithayush/go-todo-app/internal/logger"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client = nil

func InitRedis(cfg *config.Config, logr logger.Logger) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})

	if err := Client.Ping(context.Background()).Err(); err != nil {
		return err
	}

	logr.Info("Connected to Redis", logger.Field("uri", cfg.RedisURI))
	return nil
}
