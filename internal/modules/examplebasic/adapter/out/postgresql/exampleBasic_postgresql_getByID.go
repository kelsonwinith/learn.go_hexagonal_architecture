package postgresql

import (
	context "context"
	errors "errors"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleBasicMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/adapter/out/postgresql/mapper"
	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicPostgresqlGetByID struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicPostgresqlGetByID(p *sharedPostgresql.Postgresql) *ExampleBasicPostgresqlGetByID {
	return &ExampleBasicPostgresqlGetByID{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleBasicPostgresqlGetByID) Execute(ctx context.Context, id string) (*exampleBasicDomain.ExampleBasic, error) {
	var entity postgresqlModel.ExampleBasicModel

	err := e.GetExecutor(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exampleBasicDomain.ExampleBasicErrNotFound
		}
		return nil, err
	}

	return exampleBasicMapper.ToExampleBasicDomain(&entity), nil
}
