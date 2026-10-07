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

type ExamplePostgresqlCreateMultiple struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExamplePostgresqlCreateMultiple(p *sharedPostgresql.Postgresql) *ExamplePostgresqlCreateMultiple {
	return &ExamplePostgresqlCreateMultiple{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExamplePostgresqlCreateMultiple) Execute(ctx context.Context, examples []*exampleBasicDomain.Example) error {
	entities := exampleBasicMapper.ToExampleModels(examples)
	if err := e.GetExecutor(ctx).Create(entities).Error; err != nil {
		return err
	}

	for i := range entities {
		*examples[i] = *exampleBasicMapper.ToExampleDomain(entities[i])
	}

	return nil
}
