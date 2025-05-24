package config

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// NewRedis mengembalikan koneksi ke Redis
func NewRedis(logger *zap.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_ADDR", "localhost:6379"),
		Password: GetEnv("REDIS_PASS", ""),
		DB:       GetEnvInt("REDIS_DB", 0),
	})

	// Test koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		logger.Fatal("Failed to connect to Redis", zap.Error(err))
	}

	return rdb
}
