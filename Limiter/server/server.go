package server

import (
	"net/http"

	"github.com/Prateek-Pandey-96/config"
	"github.com/Prateek-Pandey-96/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(dependency *config.Dependency) {
	dependency.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	dependency.Router.Any("/limit/*proxyPath",
		middleware.LimitingMiddleware(dependency),
		// middleware.GetResponse(dependency))
	)

}

func StartServer(dependency *config.Dependency) {
	initializeRoutes(dependency)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: dependency.Router,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
