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

type ExamplePostgresqlDelete struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExamplePostgresqlDelete(p *sharedPostgresql.Postgresql) *ExamplePostgresqlDelete {
	return &ExamplePostgresqlDelete{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExamplePostgresqlDelete) Execute(ctx context.Context, id string, deletedBy int64) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleModel{}).
		Where("id = ?", id).
		Update("deleted_by", deletedBy)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exampleBasicDomain.ExampleErrNotFound
	}

	return e.GetExecutor(ctx).Where("id = ?", id).Delete(&postgresqlModel.ExampleModel{}).Error
}
