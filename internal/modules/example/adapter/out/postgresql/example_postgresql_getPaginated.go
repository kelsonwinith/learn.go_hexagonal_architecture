package postgresql

import (
	context "context"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/out/postgresql/mapper"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExamplePostgresqlGetPaginated struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExamplePostgresqlGetPaginated(p *sharedPostgresql.Postgresql) *ExamplePostgresqlGetPaginated {
	return &ExamplePostgresqlGetPaginated{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExamplePostgresqlGetPaginated) Execute(ctx context.Context, limit, offset int) ([]*exampleDomain.Example, int64, error) {
	var total int64
	if err := e.GetExecutor(ctx).Model(&postgresqlModel.ExampleModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var entities []*postgresqlModel.ExampleModel
	if err := e.GetExecutor(ctx).Order("created_at DESC").Limit(limit).Offset(offset).Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return exampleMapper.ToExampleDomains(entities), total, nil
}
