package middleware

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Prateek-Pandey-96/config"
	"github.com/gin-gonic/gin"
)

func GetResponse(dependency *config.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		targetBasePath := os.Getenv("TARGET")
		proxyPath := ctx.Param("proxyPath")
		targetUrl, err := url.Parse(targetBasePath + proxyPath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "target URL error"})
		}

		targetUrl.RawQuery = ctx.Request.URL.RawQuery

		req, err := http.NewRequest(ctx.Request.Method, targetUrl.String(), ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "request creation failed"})
		}
		defer req.Body.Close()
		req.Header = ctx.Request.Header.Clone()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": "request forwarding failed"})
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				ctx.Writer.Header().Set(key, value)
			}
		}

		if _, err := io.Copy(ctx.Writer, resp.Body); err != nil {
			log.Println("Error copying response body:", err)
		}
	}
}
