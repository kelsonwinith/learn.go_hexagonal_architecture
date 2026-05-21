package postgresql

import (
	context "context"
	time "time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExamplePostgresqlGetAll struct {
	*postgresql.Postgresql
}

func NewExamplePostgresqlGetAll(p *postgresql.Postgresql) *ExamplePostgresqlGetAll {
	return &ExamplePostgresqlGetAll{Postgresql: p}
}

func (e *ExamplePostgresqlGetAll) Execute(ctx context.Context) ([]*exampleDomain.Example, error) {
	var dtos []exampleGetAllDTO
	query := `SELECT id, name, description, created_at, updated_at FROM examples ORDER BY created_at DESC`

	err := e.GetExecutor(ctx).SelectContext(ctx, &dtos, query)
	if err != nil {
		return nil, err
	}

	examples := make([]*exampleDomain.Example, len(dtos))
	for i, dto := range dtos {
		examples[i] = dto.toDomain()
	}

	return examples, nil
}

type exampleGetAllDTO struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (d *exampleGetAllDTO) toDomain() *exampleDomain.Example {
	return &exampleDomain.Example{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
