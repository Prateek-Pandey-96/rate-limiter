package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func GetRedisClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	return client
}
