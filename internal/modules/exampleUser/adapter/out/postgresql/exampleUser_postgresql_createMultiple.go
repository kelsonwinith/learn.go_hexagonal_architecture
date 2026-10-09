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

type ExampleUserPostgresqlCreateMultiple struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPostgresqlCreateMultiple(p *sharedPostgresql.Postgresql) *ExampleUserPostgresqlCreateMultiple {
	return &ExampleUserPostgresqlCreateMultiple{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUserPostgresqlCreateMultiple) Execute(ctx context.Context, examples []*exampleUserDomain.ExampleUser) error {
	entities := exampleUserMapper.ToExampleUserModels(examples)
	if err := e.GetExecutor(ctx).Create(entities).Error; err != nil {
		return err
	}

	for i := range entities {
		*examples[i] = *exampleUserMapper.ToExampleUserDomain(entities[i])
	}

	return nil
}
