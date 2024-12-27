package limiter

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Limiter interface {
	Init(ctx context.Context, client *redis.Client)
	IsRateLimited(domain string, limit int) (bool, error)
}
