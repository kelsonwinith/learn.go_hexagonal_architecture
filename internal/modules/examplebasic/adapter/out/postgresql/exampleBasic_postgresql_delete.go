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

type ExampleBasicPostgresqlDelete struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicPostgresqlDelete(p *sharedPostgresql.Postgresql) *ExampleBasicPostgresqlDelete {
	return &ExampleBasicPostgresqlDelete{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleBasicPostgresqlDelete) Execute(ctx context.Context, id string, deletedBy int64) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleBasicModel{}).
		Where("id = ?", id).
		Update("deleted_by", deletedBy)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exampleBasicDomain.ExampleBasicErrNotFound
	}

	return e.GetExecutor(ctx).Where("id = ?", id).Delete(&postgresqlModel.ExampleBasicModel{}).Error
}
