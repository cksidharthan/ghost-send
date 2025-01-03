package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz - health check endpoint.
func Healthz(engine *gin.Engine) {
	engine.GET("/healthz", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})
}
