package repository

import (
	"context"
	"database/sql"

	"github.com/yudai2929/task-app/pkg/lib/errors"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type txKey struct{}

func withTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func getTx(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)
	return tx, ok
}

func getDB(ctx context.Context, db *sql.DB) DB {
	if tx, ok := getTx(ctx); ok {
		return tx
	}
	return db
}

func runInTransaction(ctx context.Context, db *sql.DB, fn func(ctx context.Context) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Convert(err)
	}

	ctx = withTx(ctx, tx)
	if err := fn(ctx); err != nil {
		tx.Rollback()
		return errors.Convert(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.Convert(err)
	}
	return nil
}
