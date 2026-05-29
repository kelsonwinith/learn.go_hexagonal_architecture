package postgresql

import (
	context "context"

	exampleMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/out/postgresql/mapper"
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
	entities := exampleMapper.ToExampleModels(examples)
	if err := e.GetExecutor(ctx).Create(entities).Error; err != nil {
		return err
	}

	for i := range entities {
		*examples[i] = *exampleMapper.ToExampleDomain(entities[i])
	}

	return nil
}
