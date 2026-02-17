package postgresql

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func (r *ExampleRepository) Create(ctx context.Context, example *domain.Example) error {
	query := `INSERT INTO examples (id, name, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, example.ID, example.Name, example.Description, example.CreatedAt, example.UpdatedAt)
	return err
}
