package postgresql

import (
	context "context"
	time "time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExamplePostgresqlUpdate struct {
	*postgresql.Postgresql
}

func NewExamplePostgresqlUpdate(p *postgresql.Postgresql) *ExamplePostgresqlUpdate {
	return &ExamplePostgresqlUpdate{Postgresql: p}
}

func (e *ExamplePostgresqlUpdate) Execute(ctx context.Context, example *exampleDomain.Example) error {
	dto := fromExampleUpdateDomain(example)
	query := `UPDATE examples SET name = :name, description = :description, updated_at = :updated_at 
			  WHERE id = :id`

	result, err := e.GetExecutor(ctx).NamedExecContext(ctx, query, dto)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return exampleDomain.ErrExampleNotFound
	}

	return nil
}

type exampleUpdateDTO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func fromExampleUpdateDomain(e *exampleDomain.Example) *exampleUpdateDTO {
	return &exampleUpdateDTO{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		UpdatedAt:   e.UpdatedAt,
	}
}
