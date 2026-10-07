package postgresql

import (
	context "context"

	gorm "gorm.io/gorm"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleBasicMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/adapter/out/postgresql/mapper"
	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicPostgresqlGetPaginated struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicPostgresqlGetPaginated(p *sharedPostgresql.Postgresql) *ExampleBasicPostgresqlGetPaginated {
	return &ExampleBasicPostgresqlGetPaginated{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleBasicPostgresqlGetPaginated) Execute(ctx context.Context, limit, offset int, search string) ([]*exampleBasicDomain.ExampleBasic, int64, error) {
	var total int64
	if err := exampleSearchScope(e.GetExecutor(ctx).Model(&postgresqlModel.ExampleBasicModel{}), search).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var entities []*postgresqlModel.ExampleBasicModel
	if err := exampleSearchScope(e.GetExecutor(ctx), search).Order("created_at DESC").Limit(limit).Offset(offset).Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return exampleBasicMapper.ToExampleBasicDomains(entities), total, nil
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
