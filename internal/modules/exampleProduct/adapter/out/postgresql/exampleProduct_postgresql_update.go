package postgresql

import (
	context "context"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductPostgresqlUpdate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductPostgresqlUpdate(p *sharedPostgresql.Postgresql) *ExampleProductPostgresqlUpdate {
	return &ExampleProductPostgresqlUpdate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleProductPostgresqlUpdate) Execute(ctx context.Context, example *exampleProductDomain.ExampleProduct) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleProductModel{}).
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
		return exampleProductDomain.ExampleProductErrNotFound
	}

	return nil
}
