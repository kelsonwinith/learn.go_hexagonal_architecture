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

type ExampleBasicPostgresqlCreateMultiple struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicPostgresqlCreateMultiple(p *sharedPostgresql.Postgresql) *ExampleBasicPostgresqlCreateMultiple {
	return &ExampleBasicPostgresqlCreateMultiple{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleBasicPostgresqlCreateMultiple) Execute(ctx context.Context, examples []*exampleBasicDomain.ExampleBasic) error {
	entities := exampleBasicMapper.ToExampleBasicModels(examples)
	if err := e.GetExecutor(ctx).Create(entities).Error; err != nil {
		return err
	}

	for i := range entities {
		*examples[i] = *exampleBasicMapper.ToExampleBasicDomain(entities[i])
	}

	return nil
}
