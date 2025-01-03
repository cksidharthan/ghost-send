package daemon

import (
	"context"
	"time"

	db "github.com/cksidharthan/share-secret/db/sqlc"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Janitor struct {
	dbStore *db.Store
	logger  *zap.SugaredLogger
	done    chan struct{}
}

func RunJanitor(lc fx.Lifecycle, dbStore *db.Store, logger *zap.SugaredLogger) {
	j := &Janitor{
		dbStore: dbStore,
		logger:  logger,
		done:    make(chan struct{}),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			j.logger.Info("starting secret janitor")
			go j.startCleanup()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			j.logger.Info("stopping secret janitor")
			close(j.done)
			return nil
		},
	})
}

// startCleanup - starts the secret janitor.
func (j *Janitor) startCleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-j.done:
			j.logger.Info("received shutdown signal, stopping secret janitor")
			return
		case <-ticker.C:
			j.logger.Info("cleaning expired secrets")
			err := j.dbStore.DeleteExpiredSecrets(context.Background())
			if err != nil {
				j.logger.Error("unable to delete expired secrets", zap.Error(err))
				continue
			}
			j.logger.Info("successfully cleaned expired secrets")
		}
	}
}
