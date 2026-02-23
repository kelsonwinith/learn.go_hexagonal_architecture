package postgresql

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExampleDelete struct {
	*postgresql.Postgresql
}

func NewExampleDelete(p *postgresql.Postgresql) *ExampleDelete {
	return &ExampleDelete{Postgresql: p}
}

func (e *ExampleDelete) Execute(ctx context.Context, id string) error {
	query := `DELETE FROM examples WHERE id = $1`

	result, err := e.DB.ExecContext(ctx, query, id)
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
