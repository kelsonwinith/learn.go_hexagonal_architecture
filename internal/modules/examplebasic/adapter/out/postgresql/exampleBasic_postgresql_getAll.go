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

type ExampleBasicPostgresqlGetAll struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicPostgresqlGetAll(p *sharedPostgresql.Postgresql) *ExampleBasicPostgresqlGetAll {
	return &ExampleBasicPostgresqlGetAll{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleBasicPostgresqlGetAll) Execute(ctx context.Context) ([]*exampleBasicDomain.ExampleBasic, error) {
	var entities []*postgresqlModel.ExampleBasicModel

	if err := e.GetExecutor(ctx).Order("created_at DESC").Find(&entities).Error; err != nil {
		return nil, err
	}

	return exampleBasicMapper.ToExampleBasicDomains(entities), nil
}
