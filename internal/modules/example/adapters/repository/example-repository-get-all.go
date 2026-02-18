package repository

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleGetAllRepository struct {
	*ExampleRepository
}

func NewExampleGetAllRepository(r *ExampleRepository) *ExampleGetAllRepository {
	return &ExampleGetAllRepository{ExampleRepository: r}
}

func (r *ExampleGetAllRepository) Execute(ctx context.Context) ([]*domain.Example, error) {
	var dtos []exampleGetAllDTO
	query := `SELECT id, name, description, created_at, updated_at FROM examples ORDER BY created_at DESC`

	err := r.db.SelectContext(ctx, &dtos, query)
	if err != nil {
		return nil, err
	}

	examples := make([]*domain.Example, len(dtos))
	for i, dto := range dtos {
		examples[i] = dto.toDomain()
	}

	return examples, nil
}

type exampleGetAllDTO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (d *exampleGetAllDTO) toDomain() *domain.Example {
	return &domain.Example{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
