package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ExampleRepository struct {
	db *sqlx.DB
}

func NewExampleRepository(db *sqlx.DB) *ExampleRepository {
	return &ExampleRepository{
		db: db,
	}
}

func (r *ExampleRepository) WithTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction rollback error: %v (original error: %w)", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
