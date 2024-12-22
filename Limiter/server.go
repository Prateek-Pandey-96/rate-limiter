package main

import (
	"net/http"
	"time"

	"github.com/Prateek-Pandey-96/config"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(d *config.Dependency) {
	r := d.Engine
	cacheClient := d.CacheClient

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to the rate-limiter endpoint!")
	})

	r.GET("/verify", func(ctx *gin.Context) {
		domain := "test_domain"
		val, _ := cacheClient.IsRateLimited(domain, 60, time.Second)
		if val {
			ctx.String(http.StatusTooManyRequests, "limited")
		}
		ctx.String(http.StatusOK, "not-limited")
	})
}

func Serve(d *config.Dependency) {
	d.Engine.Run(":3131")
}
