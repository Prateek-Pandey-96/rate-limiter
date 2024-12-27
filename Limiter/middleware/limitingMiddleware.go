package middleware

import (
	"net/http"

	"github.com/Prateek-Pandey-96/config"
	"github.com/gin-gonic/gin"
)

func LimitingMiddleware(dependency *config.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		domain, ok := ctx.GetQuery("domain_key")
		if !ok {
			domain = "default_domain"
		}

		limits := dependency.Limits
		domainLimit, ok := limits[domain]
		if !ok {
			domainLimit = 50
		}

		var val bool
		if domainLimit >= 200 {
			val, _ = dependency.RollingWindowClient.IsRateLimited(domain, domainLimit)
		} else {
			val, _ = dependency.TokenBucketClient.IsRateLimited(domain, domainLimit)
		}

		if val {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, "Too many requests")
			return
		}

		ctx.Next()
	}
}
