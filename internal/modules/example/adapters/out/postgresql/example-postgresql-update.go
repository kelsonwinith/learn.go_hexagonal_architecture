package postgresql

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedpostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExampleUpdateRepository struct {
	*sharedpostgresql.Repository
}

func NewExampleUpdateRepository(r *sharedpostgresql.Repository) *ExampleUpdateRepository {
	return &ExampleUpdateRepository{Repository: r}
}

func (r *ExampleUpdateRepository) Execute(ctx context.Context, example *domain.Example) error {
	dto := fromExampleUpdateDomain(example)
	query := `UPDATE examples SET name = :name, description = :description, updated_at = :updated_at 
			  WHERE id = :id`

	result, err := r.DB.NamedExecContext(ctx, query, dto)
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

type exampleUpdateDTO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func fromExampleUpdateDomain(e *domain.Example) *exampleUpdateDTO {
	return &exampleUpdateDTO{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		UpdatedAt:   e.UpdatedAt,
	}
}
