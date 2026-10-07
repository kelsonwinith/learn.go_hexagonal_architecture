package postgresql

import (
	context "context"

	exampleBasicMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/adapter/out/postgresql/mapper"
	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExamplePostgresqlCreate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExamplePostgresqlCreate(p *sharedPostgresql.Postgresql) *ExamplePostgresqlCreate {
	return &ExamplePostgresqlCreate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExamplePostgresqlCreate) Execute(ctx context.Context, example *exampleBasicDomain.Example) error {
	entity := exampleBasicMapper.ToExampleModel(example)
	if err := e.GetExecutor(ctx).Create(entity).Error; err != nil {
		return err
	}

	*example = *exampleBasicMapper.ToExampleDomain(entity)

	return nil
}
