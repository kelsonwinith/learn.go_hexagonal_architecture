package postgresql

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExamplePostgresqlCreateMultiple struct {
	*sharedPostgresql.Postgresql
}

func NewExamplePostgresqlCreateMultiple(p *sharedPostgresql.Postgresql) *ExamplePostgresqlCreateMultiple {
	return &ExamplePostgresqlCreateMultiple{Postgresql: p}
}

func (e *ExamplePostgresqlCreateMultiple) Execute(ctx context.Context, examples []*exampleDomain.Example) error {
	return e.GetExecutor(ctx).Create(toEntities(examples)).Error
}
