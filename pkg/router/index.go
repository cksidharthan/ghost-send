package router

import (
	"bytes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// IndexPage - index page endpoint.
func IndexPage(engine *gin.Engine) {
	engine.GET("/", func(context *gin.Context) {
		html, err := os.ReadFile("templates/index.html")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read index.html file"})
			return
		}
		context.DataFromReader(http.StatusOK, int64(len(html)), "text/html", bytes.NewReader(html), nil)
	})
}

func SecretView(engine *gin.Engine) {
	engine.GET("/access/:id", func(context *gin.Context) {
		html, err := os.ReadFile("templates/secret-view.html")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read secret-view.html file"})
			return
		}
		context.DataFromReader(http.StatusOK, int64(len(html)), "text/html", bytes.NewReader(html), nil)
	})
}
