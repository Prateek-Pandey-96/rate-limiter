package main

import (
	"context"
	"time"

	"github.com/Prateek-Pandey-96/config"
	"github.com/Prateek-Pandey-96/limiter"
	"github.com/Prateek-Pandey-96/redis"
	"github.com/Prateek-Pandey-96/server"
	"github.com/gin-gonic/gin"
)

func main() {
	var rollingWindowClient limiter.Limiter = &limiter.RollingWindow{}
	var tokenBucketClient limiter.Limiter = &limiter.TokenBucket{}

	ctx := context.Background()
	redisClient := redis.GetRedisClient(ctx)

	rollingWindowClient.Init(ctx, redisClient)
	tokenBucketClient.Init(ctx, redisClient)

	dependency := &config.Dependency{
		Router:              gin.Default(),
		Limits:              pollLimits(),
		RollingWindowClient: rollingWindowClient,
		TokenBucketClient:   tokenBucketClient,
	}

	go updateLimits(dependency, 15*time.Minute)
	server.StartServer(dependency)
}

func updateLimits(dependency *config.Dependency, duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		dependency.Limits = pollLimits()
	}
}

func pollLimits() map[string]int {
	// api call can be plugged here for updating limits can be added here
	return map[string]int{
		"param_value_1": 50,
		"param_value_2": 100,
	}
}
