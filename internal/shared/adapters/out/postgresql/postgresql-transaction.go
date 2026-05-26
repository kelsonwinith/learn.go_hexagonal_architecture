package postgresql

import (
	context "context"
	sql "database/sql"
	errors "errors"
	fmt "fmt"

	sqlx "github.com/jmoiron/sqlx"
)

type PostgresqlTransaction struct {
	postgresql *Postgresql
}

func NewPostgresqlTransaction(postgresql *Postgresql) *PostgresqlTransaction {
	return &PostgresqlTransaction{postgresql: postgresql}
}

func (pt *PostgresqlTransaction) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if _, ok := ctx.Value(txKey).(*sqlx.Tx); ok {
		return fn(ctx)
	}

	tx, err := pt.postgresql.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	txCtx := context.WithValue(ctx, txKey, tx)

	if err := fn(txCtx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil && !errors.Is(rbErr, sql.ErrTxDone) {
			return fmt.Errorf("transaction rollback error: %v (original error: %w)", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
