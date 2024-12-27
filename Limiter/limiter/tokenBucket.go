package limiter

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenBucket struct {
	client *redis.Client
	ctx    context.Context
}

func (tb *TokenBucket) Init(ctx context.Context, client *redis.Client) {
	tb.ctx = ctx
	tb.client = client
}

func (tb *TokenBucket) IsRateLimited(domain string, limit int) (bool, error) {
	key := fmt.Sprintf("token_bucket:%s", domain)

	scriptContent, err := os.ReadFile("./luaScripts/tokenBucket.lua")
	if err != nil {
		return false, err
	}

	// lua script for atomicity
	result, err := tb.client.Eval(tb.ctx,
		string(scriptContent),
		[]string{key},
		limit,
		time.Now().Unix(),
		10,
	).Result()

	if err != nil {
		return false, err
	}

	return result.(int64) == 1, nil
}
