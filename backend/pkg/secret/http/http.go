package http

import (
	"github.com/cksidharthan/share-secret/pkg/secret/svc"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type SecretHandler struct {
	fx.In

	SecretsSvc svc.Service

	// the name of the struct field in the Router struct in router.go
	Routes *gin.RouterGroup `name:"baseRoutes"`
}

func New(secretsHandler SecretHandler) {
	secretsHandler.Routes.GET("/secrets/:id", getSecret(secretsHandler))
	secretsHandler.Routes.POST("/secrets", postSecret(secretsHandler))
}
