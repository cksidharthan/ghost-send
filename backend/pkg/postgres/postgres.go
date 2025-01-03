package postgres

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/cksidharthan/share-secret/db/sqlc"
	"github.com/cksidharthan/share-secret/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(lc fx.Lifecycle, envCfg *config.Config, zapLogger *zap.SugaredLogger) (*db.Store, error) {
	pgDB, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		envCfg.PostgresUser,
		envCfg.PostgresPassword,
		envCfg.PostgresHost,
		envCfg.PostgresPort,
		envCfg.PostgresDB,
		envCfg.PostgresSSLMode,
	))
	if err != nil {
		zapLogger.Error("unable to connect to database", zap.Error(err))
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	driver, err := postgres.WithInstance(pgDB, &postgres.Config{})
	if err != nil {
		zapLogger.Error("unable to get database driver", zap.Error(err))
		return nil, fmt.Errorf("unable to get database driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", envCfg.MigrationsPath),
		"postgres", driver)
	if err != nil {
		zapLogger.Error("unable to get migrate instance", zap.Error(err))
		return nil, fmt.Errorf("unable to get migrate instance: %w", err)
	}

	zapLogger.Info("running migrations")

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			zapLogger.Info("no migrations to run")
		} else {
			zapLogger.Error("unable to run migrations", zap.Error(err))
			return nil, fmt.Errorf("unable to run migrations: %w", err)
		}
	}

	zapLogger.Info("migrations completed")

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			zapLogger.Info("closing db connection")
			return pgDB.Close()
		},
	})

	return db.NewStore(pgDB, zapLogger), nil
}
