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

type ExampleBasicPostgresqlUpdate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicPostgresqlUpdate(p *sharedPostgresql.Postgresql) *ExampleBasicPostgresqlUpdate {
	return &ExampleBasicPostgresqlUpdate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleBasicPostgresqlUpdate) Execute(ctx context.Context, example *exampleBasicDomain.ExampleBasic) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleBasicModel{}).
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
		return exampleBasicDomain.ExampleBasicErrNotFound
	}

	return nil
}
