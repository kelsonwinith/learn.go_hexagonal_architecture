package postgresql

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExampleCreateMultiple struct {
	*postgresql.Postgresql
}

func NewExampleCreateMultiple(p *postgresql.Postgresql) *ExampleCreateMultiple {
	return &ExampleCreateMultiple{Postgresql: p}
}

func (e *ExampleCreateMultiple) Execute(ctx context.Context, examples []*exampleDomain.Example) error {
	query := `INSERT INTO examples (id, name, description, created_at, updated_at) 
			  VALUES (:id, :name, :description, :created_at, :updated_at)`

	_, err := e.GetExecutor(ctx).NamedExecContext(ctx, query, examples)
	return err
}
