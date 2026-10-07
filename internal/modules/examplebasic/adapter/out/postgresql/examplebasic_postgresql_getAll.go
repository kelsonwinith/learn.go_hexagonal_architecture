package postgresql

import (
	context "context"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleBasicMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/adapter/out/postgresql/mapper"
	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExamplePostgresqlGetAll struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExamplePostgresqlGetAll(p *sharedPostgresql.Postgresql) *ExamplePostgresqlGetAll {
	return &ExamplePostgresqlGetAll{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExamplePostgresqlGetAll) Execute(ctx context.Context) ([]*exampleBasicDomain.Example, error) {
	var entities []*postgresqlModel.ExampleModel

	if err := e.GetExecutor(ctx).Order("created_at DESC").Find(&entities).Error; err != nil {
		return nil, err
	}

	return exampleBasicMapper.ToExampleDomains(entities), nil
}
