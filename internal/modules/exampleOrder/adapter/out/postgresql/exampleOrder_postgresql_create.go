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

type ExampleOrderPostgresqlCreate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderPostgresqlCreate(p *sharedPostgresql.Postgresql) *ExampleOrderPostgresqlCreate {
	return &ExampleOrderPostgresqlCreate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleOrderPostgresqlCreate) Execute(ctx context.Context, exampleOrder *exampleOrderDomain.ExampleOrder) error {
	entity := exampleOrderMapper.ToExampleOrderModel(exampleOrder)
	if err := e.GetExecutor(ctx).Create(entity).Error; err != nil {
		return err
	}

	exampleOrder.ID = entity.ID
	exampleOrder.CreatedAt = entity.CreatedAt
	exampleOrder.UpdatedAt = entity.UpdatedAt

	return nil
}
