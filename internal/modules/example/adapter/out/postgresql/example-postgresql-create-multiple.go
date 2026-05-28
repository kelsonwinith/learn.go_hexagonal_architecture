package postgresql

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

type ExamplePostgresqlCreateMultiple struct {
	*sharedPostgresql.Postgresql
}

func NewExamplePostgresqlCreateMultiple(p *sharedPostgresql.Postgresql) *ExamplePostgresqlCreateMultiple {
	return &ExamplePostgresqlCreateMultiple{Postgresql: p}
}

func (e *ExamplePostgresqlCreateMultiple) Execute(ctx context.Context, examples []*exampleDomain.Example) error {
	entities := toExampleModels(examples)
	if err := e.GetExecutor(ctx).Create(entities).Error; err != nil {
		return err
	}

	for i := range entities {
		*examples[i] = *toExampleDomain(entities[i])
	}

	return nil
}
