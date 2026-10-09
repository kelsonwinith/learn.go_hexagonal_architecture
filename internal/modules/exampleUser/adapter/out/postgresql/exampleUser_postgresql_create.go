package postgresql

import (
	context "context"

	exampleUserMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/out/postgresql/mapper"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserPostgresqlCreate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPostgresqlCreate(p *sharedPostgresql.Postgresql) *ExampleUserPostgresqlCreate {
	return &ExampleUserPostgresqlCreate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUserPostgresqlCreate) Execute(ctx context.Context, example *exampleUserDomain.ExampleUser) error {
	entity := exampleUserMapper.ToExampleUserModel(example)
	if err := e.GetExecutor(ctx).Create(entity).Error; err != nil {
		return err
	}

	*example = *exampleUserMapper.ToExampleUserDomain(entity)

	return nil
}
