package middleware

import (
	"net/http"
	"os"

	"github.com/Prateek-Pandey-96/config"
	"github.com/gin-gonic/gin"
)

func LimitingMiddleware(dependency *config.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limitingParam := os.Getenv("PARAM")
		domain, ok := ctx.GetQuery(limitingParam)
		if !ok {
			domain = "default_param"
		}

		limits := dependency.Limits
		domainLimit, ok := limits[domain]
		if !ok {
			domainLimit = config.DEFAULT_LIMIT
		}

		var val bool
		if domainLimit < config.ALGO_SWITCH_LIMIT {
			val, _ = dependency.RollingWindowClient.IsRateLimited(domain, domainLimit)
		} else {
			val, _ = dependency.TokenBucketClient.IsRateLimited(domain, domainLimit)
		}

		if val {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, config.TOO_MANY_REQUESTS)
			return
		}

		ctx.Next()
	}
}
