package repository

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleGetByIDRepository struct {
	*Repository
}

func NewExampleGetByIDRepository(r *Repository) *ExampleGetByIDRepository {
	return &ExampleGetByIDRepository{Repository: r}
}

func (r *ExampleGetByIDRepository) Execute(ctx context.Context, id string) (*domain.Example, error) {
	var dto exampleGetByIDDTO
	query := `SELECT id, name, description, created_at, updated_at FROM examples WHERE id = $1`

	err := r.db.GetContext(ctx, &dto, query, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, domain.ErrExampleNotFound
		}
		return nil, err
	}

	return dto.toDomain(), nil
}

type exampleGetByIDDTO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (d *exampleGetByIDDTO) toDomain() *domain.Example {
	return &domain.Example{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
