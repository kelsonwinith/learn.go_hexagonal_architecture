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

type ExampleAdvancedChildPostgresqlCreateMultiple struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleAdvancedChildPostgresqlCreateMultiple(p *sharedPostgresql.Postgresql) *ExampleAdvancedChildPostgresqlCreateMultiple {
	return &ExampleAdvancedChildPostgresqlCreateMultiple{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleAdvancedChildPostgresqlCreateMultiple) Execute(ctx context.Context, children []*exampleAdvancedDomain.Child) error {
	entities := exampleAdvancedMapper.ToChildModels(children)
	if err := e.GetExecutor(ctx).Create(entities).Error; err != nil {
		return err
	}

	for i := range entities {
		children[i].ID = entities[i].ID
		children[i].CreatedAt = entities[i].CreatedAt
		children[i].UpdatedAt = entities[i].UpdatedAt
	}

	return nil
}
