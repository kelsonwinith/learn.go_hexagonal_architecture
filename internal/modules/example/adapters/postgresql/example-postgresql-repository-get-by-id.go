package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func (r *ExampleRepository) GetByID(ctx context.Context, id string) (*domain.Example, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM examples WHERE id = $1`
	var example domain.Example
	err := r.db.GetContext(ctx, &example, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrExampleNotFound
		}
		return nil, err
	}
	return &example, nil
}
