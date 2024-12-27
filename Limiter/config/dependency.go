package config

import (
	"github.com/Prateek-Pandey-96/limiter"
	"github.com/gin-gonic/gin"
)

type Dependency struct {
	Router              *gin.Engine
	Limits              map[string]int
	RollingWindowClient limiter.Limiter
	TokenBucketClient   limiter.Limiter
}
