package postgresql

import (
	context "context"
	errors "errors"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleOrderMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/postgresql/mapper"
	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderPostgresqlGetByID struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderPostgresqlGetByID(p *sharedPostgresql.Postgresql) *ExampleOrderPostgresqlGetByID {
	return &ExampleOrderPostgresqlGetByID{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleOrderPostgresqlGetByID) Execute(ctx context.Context, id string) (*exampleOrderDomain.ExampleOrder, error) {
	var entity postgresqlModel.ExampleOrderModel

	err := e.GetExecutor(ctx).Preload("Products").Where("id = ?", id).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exampleOrderDomain.ExampleOrderErrNotFound
		}
		return nil, err
	}

	return exampleOrderMapper.ToExampleOrderDomain(&entity), nil
}
