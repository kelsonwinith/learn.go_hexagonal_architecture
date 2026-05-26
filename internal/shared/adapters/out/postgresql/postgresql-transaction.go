package postgresql

import (
	context "context"
	fmt "fmt"

	gorm "gorm.io/gorm"
)

type PostgresqlTransaction struct {
	postgresql *Postgresql
}

func NewPostgresqlTransaction(postgresql *Postgresql) *PostgresqlTransaction {
	return &PostgresqlTransaction{postgresql: postgresql}
}

func (pt *PostgresqlTransaction) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if _, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return fn(ctx)
	}

	err := pt.postgresql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, txKey, tx)
		return fn(txCtx)
	})
	if err != nil {
		return fmt.Errorf("transaction error: %w", err)
	}

	return nil
}
