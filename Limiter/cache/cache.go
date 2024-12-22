package cache

import (
	"context"
	"time"
)

type ICache interface {
	Init(ctx context.Context)
	IsRateLimited(domain string, limit int64, duration time.Duration) (bool, error)
}
