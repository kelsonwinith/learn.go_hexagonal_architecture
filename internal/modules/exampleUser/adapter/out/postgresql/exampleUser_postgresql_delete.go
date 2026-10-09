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

type ExampleUserPostgresqlDelete struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPostgresqlDelete(p *sharedPostgresql.Postgresql) *ExampleUserPostgresqlDelete {
	return &ExampleUserPostgresqlDelete{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUserPostgresqlDelete) Execute(ctx context.Context, id string, deletedBy int64) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleUserModel{}).
		Where("id = ?", id).
		Update("deleted_by", deletedBy)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exampleUserDomain.ExampleUserErrNotFound
	}

	return e.GetExecutor(ctx).Where("id = ?", id).Delete(&postgresqlModel.ExampleUserModel{}).Error
}
