package postgresql

import (
	context "context"
	errors "errors"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleAdvancedMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/adapter/out/postgresql/mapper"
	exampleAdvancedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Types
// ============================================================================

type ExampleAdvancedPostgresqlGetByID struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleAdvancedPostgresqlGetByID(p *sharedPostgresql.Postgresql) *ExampleAdvancedPostgresqlGetByID {
	return &ExampleAdvancedPostgresqlGetByID{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleAdvancedPostgresqlGetByID) Execute(ctx context.Context, id string) (*exampleAdvancedDomain.Parent, error) {
	var entity postgresqlModel.ExampleAdvancedParentModel

	err := e.GetExecutor(ctx).Preload("Children").Where("id = ?", id).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exampleAdvancedDomain.ExampleAdvancedErrNotFound
		}
		return nil, err
	}

	return exampleAdvancedMapper.ToParentDomain(&entity), nil
}
