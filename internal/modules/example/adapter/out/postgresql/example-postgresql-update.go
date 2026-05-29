package postgresql

import (
	context "context"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

type ExamplePostgresqlUpdate struct {
	*sharedPostgresql.Postgresql
}

func NewExamplePostgresqlUpdate(p *sharedPostgresql.Postgresql) *ExamplePostgresqlUpdate {
	return &ExamplePostgresqlUpdate{Postgresql: p}
}

func (e *ExamplePostgresqlUpdate) Execute(ctx context.Context, example *exampleDomain.Example) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleModel{}).
		Where("id = ?", example.ID).
		Updates(map[string]interface{}{
			"name":        example.Name,
			"description": example.Description,
			"updated_at":  example.UpdatedAt,
		})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exampleDomain.ExampleErrNotFound
	}

	return nil
}
