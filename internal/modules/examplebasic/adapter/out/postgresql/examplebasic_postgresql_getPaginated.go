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

type ExamplePostgresqlGetPaginated struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExamplePostgresqlGetPaginated(p *sharedPostgresql.Postgresql) *ExamplePostgresqlGetPaginated {
	return &ExamplePostgresqlGetPaginated{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExamplePostgresqlGetPaginated) Execute(ctx context.Context, limit, offset int, search string) ([]*exampleBasicDomain.Example, int64, error) {
	var total int64
	if err := exampleSearchScope(e.GetExecutor(ctx).Model(&postgresqlModel.ExampleModel{}), search).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var entities []*postgresqlModel.ExampleModel
	if err := exampleSearchScope(e.GetExecutor(ctx), search).Order("created_at DESC").Limit(limit).Offset(offset).Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return exampleBasicMapper.ToExampleDomains(entities), total, nil
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
