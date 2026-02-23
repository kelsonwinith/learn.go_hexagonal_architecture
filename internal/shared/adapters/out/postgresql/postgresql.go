package postgresql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Postgresql struct {
	DB *sqlx.DB
}

func NewPostgresql(db *sqlx.DB) *Postgresql {
	return &Postgresql{
		DB: db,
	}
}

func (p *Postgresql) WithTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := p.DB.BeginTxx(ctx, nil)
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
