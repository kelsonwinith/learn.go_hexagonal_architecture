package postgresql

import (
	context "context"

	gorm "gorm.io/gorm"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleOrderMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/postgresql/mapper"
	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderPostgresqlGetPaginated struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderPostgresqlGetPaginated(p *sharedPostgresql.Postgresql) *ExampleOrderPostgresqlGetPaginated {
	return &ExampleOrderPostgresqlGetPaginated{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleOrderPostgresqlGetPaginated) Execute(ctx context.Context, limit, offset int, search string) ([]*exampleOrderDomain.ExampleOrder, int64, error) {
	var total int64
	if err := exampleOrderSearchScope(e.GetExecutor(ctx).Model(&postgresqlModel.ExampleOrderModel{}), search).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var entities []*postgresqlModel.ExampleOrderModel
	if err := exampleOrderSearchScope(e.GetExecutor(ctx), search).
		Preload("Products").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return exampleOrderMapper.ToExampleOrderDomains(entities), total, nil
}

// ============================================================================
// Functions
// ============================================================================

func exampleOrderSearchScope(db *gorm.DB, search string) *gorm.DB {
	if search == "" {
		return db
	}

	pattern := "%" + search + "%"
	return db.Where("name ILIKE ? OR description ILIKE ?", pattern, pattern)
}
