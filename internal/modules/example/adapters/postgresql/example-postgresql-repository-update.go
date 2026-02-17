package postgresql

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func (r *ExampleRepository) Update(ctx context.Context, example *domain.Example) error {
	query := `UPDATE examples SET name = $1, description = $2, updated_at = $3 WHERE id = $4`
	result, err := r.db.ExecContext(ctx, query, example.Name, example.Description, example.UpdatedAt, example.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrExampleNotFound
	}
	return nil
}
