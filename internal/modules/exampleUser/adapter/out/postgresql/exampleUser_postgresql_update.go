package postgresql

import (
	context "context"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserPostgresqlUpdate struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPostgresqlUpdate(p *sharedPostgresql.Postgresql) *ExampleUserPostgresqlUpdate {
	return &ExampleUserPostgresqlUpdate{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUserPostgresqlUpdate) Execute(ctx context.Context, example *exampleUserDomain.ExampleUser) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleUserModel{}).
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
		return exampleUserDomain.ExampleUserErrNotFound
	}

	return nil
}
