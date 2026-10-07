package postgresql

import (
	context "context"

	exampleAdvancedMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/adapter/out/postgresql/mapper"
	exampleAdvancedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleAdvancedPostgresqlCreate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleAdvancedPostgresqlCreate(p *sharedPostgresql.Postgresql) *ExampleAdvancedPostgresqlCreate {
	return &ExampleAdvancedPostgresqlCreate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleAdvancedPostgresqlCreate) Execute(ctx context.Context, parent *exampleAdvancedDomain.Parent) error {
	entity := exampleAdvancedMapper.ToParentModel(parent)
	if err := e.GetExecutor(ctx).Create(entity).Error; err != nil {
		return err
	}

	parent.ID = entity.ID
	parent.CreatedAt = entity.CreatedAt
	parent.UpdatedAt = entity.UpdatedAt

	return nil
}
