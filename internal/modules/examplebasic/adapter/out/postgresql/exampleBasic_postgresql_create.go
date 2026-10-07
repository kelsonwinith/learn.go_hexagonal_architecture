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

type ExampleBasicPostgresqlCreate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicPostgresqlCreate(p *sharedPostgresql.Postgresql) *ExampleBasicPostgresqlCreate {
	return &ExampleBasicPostgresqlCreate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleBasicPostgresqlCreate) Execute(ctx context.Context, example *exampleBasicDomain.ExampleBasic) error {
	entity := exampleBasicMapper.ToExampleBasicModel(example)
	if err := e.GetExecutor(ctx).Create(entity).Error; err != nil {
		return err
	}

	*example = *exampleBasicMapper.ToExampleBasicDomain(entity)

	return nil
}
