package cmd

import (
	"github.com/cksidharthan/share-secret/pkg/config"
	"github.com/cksidharthan/share-secret/pkg/daemon"
	"github.com/cksidharthan/share-secret/pkg/logger"
	"github.com/cksidharthan/share-secret/pkg/postgres"
	"github.com/cksidharthan/share-secret/pkg/router"
	secretHttp "github.com/cksidharthan/share-secret/pkg/secret/http"
	secretSvc "github.com/cksidharthan/share-secret/pkg/secret/svc"
	"go.uber.org/fx"
)

func Start() {
	app := fx.New(
		fx.Provide(
			config.New,
			logger.New,
			router.New,
			postgres.New,
			secretSvc.New,
		),
		fx.Invoke(
			router.Healthz,
			daemon.RunJanitor,
			secretHttp.New,
		),
	)

	app.Run()
}
