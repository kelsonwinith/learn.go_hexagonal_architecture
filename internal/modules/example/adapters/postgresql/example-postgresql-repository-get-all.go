package postgresql

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func (r *ExampleRepository) GetAll(ctx context.Context) ([]*domain.Example, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM examples`
	var examples []*domain.Example
	err := r.db.SelectContext(ctx, &examples, query)
	return examples, err
}
