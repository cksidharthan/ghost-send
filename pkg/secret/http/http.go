package http

import (
	"github.com/cksidharthan/ghost-send/pkg/secret/svc"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type SecretHandler struct {
	fx.In

	SecretsSvc svc.Service
	Routes     *gin.RouterGroup `name:"apiV1Routes"`
}

func New(secretsHandler SecretHandler) {
	secretsHandler.Routes.POST("/secrets/:id", getSecret(secretsHandler))
	secretsHandler.Routes.POST("/secrets", postSecret(secretsHandler))
	secretsHandler.Routes.GET("/secrets/:id/status", getSecretStatus(secretsHandler))
}
