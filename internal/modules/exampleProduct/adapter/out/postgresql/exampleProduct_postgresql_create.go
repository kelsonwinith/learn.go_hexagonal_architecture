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

type ExampleProductPostgresqlCreate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductPostgresqlCreate(p *sharedPostgresql.Postgresql) *ExampleProductPostgresqlCreate {
	return &ExampleProductPostgresqlCreate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleProductPostgresqlCreate) Execute(ctx context.Context, example *exampleProductDomain.ExampleProduct) error {
	entity := exampleProductMapper.ToExampleProductModel(example)
	if err := e.GetExecutor(ctx).Create(entity).Error; err != nil {
		return err
	}

	*example = *exampleProductMapper.ToExampleProductDomain(entity)

	return nil
}
