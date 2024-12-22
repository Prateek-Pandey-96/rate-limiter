package config

import (
	"github.com/Prateek-Pandey-96/cache"
	"github.com/gin-gonic/gin"
)

type Dependency struct {
	Engine      *gin.Engine
	CacheClient cache.ICache
}
