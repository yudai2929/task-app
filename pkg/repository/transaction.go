package repository

import (
	"context"
	"database/sql"

	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
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

func runInTransaction(ctx context.Context, db *sql.DB, fn func(ctx context.Context) error) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Convert(err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = errors.Newf(codes.CodeInternal, "panic recovered: %v", r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				err = errors.Convert(err)
			}
		}
	}()

	ctx = withTx(ctx, tx)
	err = fn(ctx)
	return err
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (t *transactionRepository) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	return runInTransaction(ctx, t.db, fn)
}
