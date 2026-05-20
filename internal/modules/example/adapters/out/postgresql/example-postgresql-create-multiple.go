package postgresql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExampleCreateMultiple struct {
	*postgresql.Postgresql
}

func NewExampleCreateMultiple(p *postgresql.Postgresql) *ExampleCreateMultiple {
	return &ExampleCreateMultiple{Postgresql: p}
}

func (e *ExampleCreateMultiple) Execute(ctx context.Context, db sqlx.ExtContext, examples []*domain.Example) error {
	query := `INSERT INTO examples (id, name, description, created_at, updated_at)
			  VALUES (:id, :name, :description, :created_at, :updated_at)`

	_, err := sqlx.NamedExecContext(ctx, db, query, examples)
	return err
}
