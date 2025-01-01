package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Store struct {
	*Queries
	db     *sql.DB
	zapLog *zap.SugaredLogger
}

func NewStore(pgDB *sql.DB, zapLog *zap.SugaredLogger) *Store {
	return &Store{
		db:      pgDB,
		Queries: New(pgDB),
		zapLog:  zapLog,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		store.zapLog.Error("unable to begin tx", zap.Error(err))
		return err
	}
	queries := New(tx)
	err = fn(queries)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			store.zapLog.Error("unable to rollback tx", zap.Error(rbErr))
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}
