package postgresql

import (
	context "context"

	exampleOrderMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/postgresql/mapper"
	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderProductPostgresqlCreateMultiple struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderProductPostgresqlCreateMultiple(p *sharedPostgresql.Postgresql) *ExampleOrderProductPostgresqlCreateMultiple {
	return &ExampleOrderProductPostgresqlCreateMultiple{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleOrderProductPostgresqlCreateMultiple) Execute(ctx context.Context, exampleOrderProducts []*exampleOrderDomain.ExampleOrderProduct) error {
	entities := exampleOrderMapper.ToExampleOrderProductModels(exampleOrderProducts)
	if err := e.GetExecutor(ctx).Create(entities).Error; err != nil {
		return err
	}

	for i := range entities {
		exampleOrderProducts[i].ID = entities[i].ID
		exampleOrderProducts[i].CreatedAt = entities[i].CreatedAt
		exampleOrderProducts[i].UpdatedAt = entities[i].UpdatedAt
	}

	return nil
}
