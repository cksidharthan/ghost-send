package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexPage - index page endpoint.
func IndexPage(engine *gin.Engine) {
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
}
