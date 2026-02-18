package repository

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleCreateRepository struct {
	*ExampleRepository
}

func NewExampleCreateRepository(r *ExampleRepository) *ExampleCreateRepository {
	return &ExampleCreateRepository{ExampleRepository: r}
}

func (r *ExampleCreateRepository) Execute(ctx context.Context, example *domain.Example) error {
	dto := fromExampleCreateDomain(example)
	query := `INSERT INTO examples (id, name, description, created_at, updated_at) 
			  VALUES (:id, :name, :description, :created_at, :updated_at)`

	_, err := r.db.NamedExecContext(ctx, query, dto)
	return err
}

type exampleCreateDTO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func fromExampleCreateDomain(e *domain.Example) *exampleCreateDTO {
	return &exampleCreateDTO{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
