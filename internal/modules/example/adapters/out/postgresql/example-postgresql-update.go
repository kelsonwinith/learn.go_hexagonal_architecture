package postgresql

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExampleUpdate struct {
	*postgresql.Postgresql
}

func NewExampleUpdate(p *postgresql.Postgresql) *ExampleUpdate {
	return &ExampleUpdate{Postgresql: p}
}

func (e *ExampleUpdate) Execute(ctx context.Context, example *domain.Example) error {
	dto := fromExampleUpdateDomain(example)
	query := `UPDATE examples SET name = :name, description = :description, updated_at = :updated_at 
			  WHERE id = :id`

	result, err := e.DB.NamedExecContext(ctx, query, dto)
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
