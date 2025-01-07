package logger

import (
	"context"
	"log"

	"github.com/cksidharthan/ghost-send/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(lc fx.Lifecycle, envCfg *config.Config) *zap.SugaredLogger {
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(envCfg.LogLevel)); err != nil {
		log.Fatalf("failed to parse log level: %v", err)
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Info("closing logger")
			return nil
		},
	})

	return logger.Sugar()
}
