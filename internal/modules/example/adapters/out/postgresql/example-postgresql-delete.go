package postgresql

import (
	context "context"

	postgresqlEntity "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/entity"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type ExamplePostgresqlDelete struct {
	*sharedPostgresql.Postgresql
}

func NewExamplePostgresqlDelete(p *sharedPostgresql.Postgresql) *ExamplePostgresqlDelete {
	return &ExamplePostgresqlDelete{Postgresql: p}
}

func (e *ExamplePostgresqlDelete) Execute(ctx context.Context, id string) error {
	result := e.GetExecutor(ctx).Where("id = ?", id).Delete(&postgresqlEntity.Example{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exampleDomain.ExampleErrNotFound
	}

	return nil
}
