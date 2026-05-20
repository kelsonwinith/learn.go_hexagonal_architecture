package postgresql

import (
	context "context"
	sql "database/sql"
	fmt "fmt"

	sqlx "github.com/jmoiron/sqlx"
)

type Ext interface {
	sqlx.ExtContext
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type txKeyType struct{}

var txKey = txKeyType{}

type Postgresql struct {
	DB *sqlx.DB
}

func NewPostgresql(db *sqlx.DB) *Postgresql {
	return &Postgresql{
		DB: db,
	}
}

func (p *Postgresql) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if _, ok := ctx.Value(txKey).(*sqlx.Tx); ok {
		return fn(ctx)
	}

	tx, err := p.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	txCtx := context.WithValue(ctx, txKey, tx)

	if err := fn(txCtx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil && !errorsIsRollbackDone(rbErr) {
			return fmt.Errorf("transaction rollback error: %v (original error: %w)", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (p *Postgresql) GetExecutor(ctx context.Context) Ext {
	if tx, ok := ctx.Value(txKey).(*sqlx.Tx); ok {
		return tx
	}
	return p.DB
}

func errorsIsRollbackDone(err error) bool {
	return err == sql.ErrTxDone
}
