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

type ExampleProductPostgresqlDelete struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductPostgresqlDelete(p *sharedPostgresql.Postgresql) *ExampleProductPostgresqlDelete {
	return &ExampleProductPostgresqlDelete{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleProductPostgresqlDelete) Execute(ctx context.Context, id string, deletedBy int64) error {
	result := e.GetExecutor(ctx).
		Model(&postgresqlModel.ExampleProductModel{}).
		Where("id = ?", id).
		Update("deleted_by", deletedBy)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exampleProductDomain.ExampleProductErrNotFound
	}

	return e.GetExecutor(ctx).Where("id = ?", id).Delete(&postgresqlModel.ExampleProductModel{}).Error
}
