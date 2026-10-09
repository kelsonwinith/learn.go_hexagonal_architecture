package postgresql

import (
	context "context"

	gorm "gorm.io/gorm"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleProductMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/adapter/out/postgresql/mapper"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductPostgresqlGetPaginated struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductPostgresqlGetPaginated(p *sharedPostgresql.Postgresql) *ExampleProductPostgresqlGetPaginated {
	return &ExampleProductPostgresqlGetPaginated{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleProductPostgresqlGetPaginated) Execute(ctx context.Context, limit, offset int, search string) ([]*exampleProductDomain.ExampleProduct, int64, error) {
	var total int64
	if err := exampleSearchScope(e.GetExecutor(ctx).Model(&postgresqlModel.ExampleProductModel{}), search).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var entities []*postgresqlModel.ExampleProductModel
	if err := exampleSearchScope(e.GetExecutor(ctx), search).Order("created_at DESC").Limit(limit).Offset(offset).Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return exampleProductMapper.ToExampleProductDomains(entities), total, nil
}

// ============================================================================
// Functions
// ============================================================================

func exampleSearchScope(db *gorm.DB, search string) *gorm.DB {
	if search == "" {
		return db
	}

	pattern := "%" + search + "%"
	return db.Where("name ILIKE ? OR description ILIKE ?", pattern, pattern)
}
