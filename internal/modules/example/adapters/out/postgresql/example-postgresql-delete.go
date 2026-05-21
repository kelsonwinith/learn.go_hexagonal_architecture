package postgresql

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExamplePostgresqlDelete struct {
	*postgresql.Postgresql
}

func NewExamplePostgresqlDelete(p *postgresql.Postgresql) *ExamplePostgresqlDelete {
	return &ExamplePostgresqlDelete{Postgresql: p}
}

func (e *ExamplePostgresqlDelete) Execute(ctx context.Context, id string) error {
	query := `DELETE FROM examples WHERE id = $1`

	result, err := e.GetExecutor(ctx).ExecContext(ctx, query, id)
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
