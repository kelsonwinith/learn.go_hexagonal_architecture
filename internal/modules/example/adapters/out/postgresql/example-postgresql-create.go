package postgresql

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedpostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExampleCreateRepository struct {
	*sharedpostgresql.Repository
}

func NewExampleCreateRepository(r *sharedpostgresql.Repository) *ExampleCreateRepository {
	return &ExampleCreateRepository{Repository: r}
}

func (r *ExampleCreateRepository) Execute(ctx context.Context, example *domain.Example) error {
	dto := fromExampleCreateDomain(example)
	query := `INSERT INTO examples (id, name, description, created_at, updated_at) 
			  VALUES (:id, :name, :description, :created_at, :updated_at)`

	_, err := r.DB.NamedExecContext(ctx, query, dto)
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
