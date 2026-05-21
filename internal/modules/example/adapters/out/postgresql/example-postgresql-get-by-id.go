package postgresql

import (
	context "context"
	sql "database/sql"
	errors "errors"
	time "time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExamplePostgresqlGetByID struct {
	*postgresql.Postgresql
}

func NewExamplePostgresqlGetByID(p *postgresql.Postgresql) *ExamplePostgresqlGetByID {
	return &ExamplePostgresqlGetByID{Postgresql: p}
}

func (e *ExamplePostgresqlGetByID) Execute(ctx context.Context, id string) (*exampleDomain.Example, error) {
	var dto exampleGetByIDDTO
	query := `SELECT id, name, description, created_at, updated_at FROM examples WHERE id = $1`

	err := e.GetExecutor(ctx).GetContext(ctx, &dto, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exampleDomain.ErrExampleNotFound
		}
		return nil, err
	}

	return dto.toDomain(), nil
}

type exampleGetByIDDTO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (d *exampleGetByIDDTO) toDomain() *exampleDomain.Example {
	return &exampleDomain.Example{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
