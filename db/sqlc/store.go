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

// AccessSecretAtomic reads and conditionally deletes a secret in a single
// serialized transaction. Concurrent callers block on the FOR UPDATE lock
// until the first transaction commits; after a single-view secret is deleted
// the second caller finds no rows and receives sql.ErrNoRows.
func (store *Store) AccessSecretAtomic(ctx context.Context, arg GetSecretByIDLockedParams) (*GetSecretByIDLockedRow, error) {
	var result GetSecretByIDLockedRow
	err := store.execTx(ctx, func(q *Queries) error {
		row, err := q.GetSecretByIDLocked(ctx, arg)
		if err != nil {
			return err
		}
		result = row

		if !result.PasswordMatches {
			return nil
		}

		if _, err := q.DecrementTries(ctx, arg.SecretID); err != nil {
			return err
		}
		if result.ShouldDelete {
			return q.DeleteSecret(ctx, arg.SecretID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}
