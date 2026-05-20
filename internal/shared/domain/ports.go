package domain

import (
	context "context"
)

type TransactionManagerPostgresql interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
