package postgresql

import (
	context "context"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleUserMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/out/postgresql/mapper"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserPostgresqlGetAll struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPostgresqlGetAll(p *sharedPostgresql.Postgresql) *ExampleUserPostgresqlGetAll {
	return &ExampleUserPostgresqlGetAll{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUserPostgresqlGetAll) Execute(ctx context.Context) ([]*exampleUserDomain.ExampleUser, error) {
	var entities []*postgresqlModel.ExampleUserModel

	if err := e.GetExecutor(ctx).Order("created_at DESC").Find(&entities).Error; err != nil {
		return nil, err
	}

	return exampleUserMapper.ToExampleUserDomains(entities), nil
}
