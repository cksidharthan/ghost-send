package cmd

import (
	"github.com/cksidharthan/ghost-send/pkg/config"
	"github.com/cksidharthan/ghost-send/pkg/daemon"
	"github.com/cksidharthan/ghost-send/pkg/logger"
	"github.com/cksidharthan/ghost-send/pkg/postgres"
	"github.com/cksidharthan/ghost-send/pkg/router"
	secretHttp "github.com/cksidharthan/ghost-send/pkg/secret/http"
	secretSvc "github.com/cksidharthan/ghost-send/pkg/secret/svc"
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
