package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func (rc *RedisClient) Init(ctx context.Context) {
	rc.ctx = ctx
	rc.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func (rc *RedisClient) IsRateLimited(domain string, limit int64, duration time.Duration) (bool, error) {
	key := fmt.Sprintf("domain_key:%s", domain)
	timestamp := time.Now().UTC().UnixMilli()

	err := rc.client.ZRemRangeByScore(rc.ctx, key, "-inf", fmt.Sprint(timestamp-int64(duration.Milliseconds()))).Err()
	if err != nil {
		return false, err
	}

	err = rc.client.ZAdd(rc.ctx, key, redis.Z{Score: float64(timestamp), Member: timestamp}).Err()
	if err != nil {
		return false, err
	}

	count, err := rc.client.ZCard(rc.ctx, key).Result()
	if err != nil {
		return false, err
	}

	return count > limit, nil
}
