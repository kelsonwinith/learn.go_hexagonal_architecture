package repository

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleDeleteRepository struct {
	*ExampleRepository
}

func NewExampleDeleteRepository(r *ExampleRepository) *ExampleDeleteRepository {
	return &ExampleDeleteRepository{ExampleRepository: r}
}

func (r *ExampleDeleteRepository) Execute(ctx context.Context, id string) error {
	query := `DELETE FROM examples WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrExampleNotFound
	}

	return nil
}
