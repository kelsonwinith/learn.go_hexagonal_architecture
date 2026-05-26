package postgresql

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExamplePostgresqlCreate struct {
	*sharedPostgresql.Postgresql
}

func NewExamplePostgresqlCreate(p *sharedPostgresql.Postgresql) *ExamplePostgresqlCreate {
	return &ExamplePostgresqlCreate{Postgresql: p}
}

func (e *ExamplePostgresqlCreate) Execute(ctx context.Context, example *exampleDomain.Example) error {
	return e.GetExecutor(ctx).Create(toEntity(example)).Error
}
