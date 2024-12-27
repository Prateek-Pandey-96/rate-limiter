package main

import (
	"context"

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
		Limits:              map[string]int{},
		RollingWindowClient: rollingWindowClient,
		TokenBucketClient:   tokenBucketClient,
	}

	server.StartServer(dependency)
}
