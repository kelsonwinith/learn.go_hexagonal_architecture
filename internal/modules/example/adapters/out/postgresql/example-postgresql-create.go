package postgresql

import (
	context "context"
	time "time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExamplePostgresqlCreate struct {
	*postgresql.Postgresql
}

func NewExamplePostgresqlCreate(p *postgresql.Postgresql) *ExamplePostgresqlCreate {
	return &ExamplePostgresqlCreate{Postgresql: p}
}

func (e *ExamplePostgresqlCreate) Execute(ctx context.Context, example *exampleDomain.Example) error {
	dto := fromExampleCreateDomain(example)
	query := `INSERT INTO examples (id, name, description, created_at, updated_at) 
			  VALUES (:id, :name, :description, :created_at, :updated_at)`

	_, err := e.GetExecutor(ctx).NamedExecContext(ctx, query, dto)
	return err
}

type exampleCreateDTO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func fromExampleCreateDomain(e *exampleDomain.Example) *exampleCreateDTO {
	return &exampleCreateDTO{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
