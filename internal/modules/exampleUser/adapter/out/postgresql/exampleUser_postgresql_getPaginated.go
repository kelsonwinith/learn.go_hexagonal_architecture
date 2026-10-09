package postgresql

import (
	context "context"

	gorm "gorm.io/gorm"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleUserMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/out/postgresql/mapper"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserPostgresqlGetPaginated struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPostgresqlGetPaginated(p *sharedPostgresql.Postgresql) *ExampleUserPostgresqlGetPaginated {
	return &ExampleUserPostgresqlGetPaginated{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUserPostgresqlGetPaginated) Execute(ctx context.Context, limit, offset int, search string) ([]*exampleUserDomain.ExampleUser, int64, error) {
	var total int64
	if err := exampleSearchScope(e.GetExecutor(ctx).Model(&postgresqlModel.ExampleUserModel{}), search).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var entities []*postgresqlModel.ExampleUserModel
	if err := exampleSearchScope(e.GetExecutor(ctx), search).Order("created_at DESC").Limit(limit).Offset(offset).Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return exampleUserMapper.ToExampleUserDomains(entities), total, nil
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
