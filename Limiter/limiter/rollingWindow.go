package limiter

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RollingWindow struct {
	client *redis.Client
	ctx    context.Context
}

func (rc *RollingWindow) Init(ctx context.Context, client *redis.Client) {
	rc.ctx = ctx
	rc.client = client
}

func (rc *RollingWindow) IsRateLimited(domain string, limit int) (bool, error) {
	key := fmt.Sprintf("domain_key:rolling_window:%s", domain)
	intervalMicroSec := 1e6

	scriptContent, err := os.ReadFile("./luaScripts/rollingWindow.lua")
	if err != nil {
		return false, err
	}

	// lua script for atomicity
	result, err := rc.client.Eval(rc.ctx,
		string(scriptContent),
		[]string{key},
		limit,
		intervalMicroSec,
		time.Now().UnixMicro(),
		time.Now().UnixNano(),
	).Result()
	if err != nil {
		return false, err
	}

	return result.(int64) == 1, nil
}
