package postgresql

import (
	context "context"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

type ExamplePostgresqlGetAll struct {
	*sharedPostgresql.Postgresql
}

func NewExamplePostgresqlGetAll(p *sharedPostgresql.Postgresql) *ExamplePostgresqlGetAll {
	return &ExamplePostgresqlGetAll{Postgresql: p}
}

func (e *ExamplePostgresqlGetAll) Execute(ctx context.Context) ([]*exampleDomain.Example, error) {
	var entities []*postgresqlModel.Example

	if err := e.GetExecutor(ctx).Order("created_at DESC").Find(&entities).Error; err != nil {
		return nil, err
	}

	return toExampleDomains(entities), nil
}
