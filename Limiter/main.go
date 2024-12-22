package main

import (
	"context"

	"github.com/Prateek-Pandey-96/cache"
	"github.com/Prateek-Pandey-96/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var cacheClient cache.ICache = &cache.RedisClient{}
	ctx := context.Background()

	cacheClient.Init(ctx)

	dependency := &config.Dependency{
		Engine:      r,
		CacheClient: cacheClient,
	}

	InitializeRoutes(dependency)
	Serve(dependency)
}
