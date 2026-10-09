package postgresql

import (
	context "context"

	exampleProductMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/adapter/out/postgresql/mapper"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductPostgresqlCreateMultiple struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductPostgresqlCreateMultiple(p *sharedPostgresql.Postgresql) *ExampleProductPostgresqlCreateMultiple {
	return &ExampleProductPostgresqlCreateMultiple{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleProductPostgresqlCreateMultiple) Execute(ctx context.Context, examples []*exampleProductDomain.ExampleProduct) error {
	entities := exampleProductMapper.ToExampleProductModels(examples)
	if err := e.GetExecutor(ctx).Create(entities).Error; err != nil {
		return err
	}

	for i := range entities {
		*examples[i] = *exampleProductMapper.ToExampleProductDomain(entities[i])
	}

	return nil
}
