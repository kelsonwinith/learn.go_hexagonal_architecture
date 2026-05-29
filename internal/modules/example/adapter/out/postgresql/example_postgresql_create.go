package postgresql

import (
	context "context"

	exampleMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/out/postgresql/mapper"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

type ExamplePostgresqlCreate struct {
	*sharedPostgresql.Postgresql
}

func NewExamplePostgresqlCreate(p *sharedPostgresql.Postgresql) *ExamplePostgresqlCreate {
	return &ExamplePostgresqlCreate{Postgresql: p}
}

func (e *ExamplePostgresqlCreate) Execute(ctx context.Context, example *exampleDomain.Example) error {
	entity := exampleMapper.ToExampleModel(example)
	if err := e.GetExecutor(ctx).Create(entity).Error; err != nil {
		return err
	}

	*example = *exampleMapper.ToExampleDomain(entity)

	return nil
}
