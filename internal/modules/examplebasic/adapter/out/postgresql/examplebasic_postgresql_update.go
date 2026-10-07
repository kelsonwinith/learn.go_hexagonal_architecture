package postgresql

import (
	context "context"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExamplePostgresqlUpdate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExamplePostgresqlUpdate(p *sharedPostgresql.Postgresql) *ExamplePostgresqlUpdate {
	return &ExamplePostgresqlUpdate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExamplePostgresqlUpdate) Execute(ctx context.Context, example *exampleBasicDomain.Example) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleModel{}).
		Where("id = ?", example.ID).
		Updates(map[string]interface{}{
			"name":        example.Name,
			"description": example.Description,
			"updated_by":  example.UpdatedBy,
			"updated_at":  example.UpdatedAt,
		})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exampleBasicDomain.ExampleErrNotFound
	}

	return nil
}
