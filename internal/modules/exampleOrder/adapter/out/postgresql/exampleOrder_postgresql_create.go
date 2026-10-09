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

func (e *ExampleOrderPostgresqlCreate) Execute(ctx context.Context, order *exampleOrderDomain.ExampleOrder) error {
	entity := exampleOrderMapper.ToExampleOrderModel(order)
	if err := e.GetExecutor(ctx).Create(entity).Error; err != nil {
		return err
	}

	order.ID = entity.ID
	order.CreatedAt = entity.CreatedAt
	order.UpdatedAt = entity.UpdatedAt

	return nil
}
