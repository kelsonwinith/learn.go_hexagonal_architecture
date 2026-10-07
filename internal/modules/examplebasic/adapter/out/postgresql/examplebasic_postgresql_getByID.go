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

type ExamplePostgresqlGetByID struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExamplePostgresqlGetByID(p *sharedPostgresql.Postgresql) *ExamplePostgresqlGetByID {
	return &ExamplePostgresqlGetByID{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExamplePostgresqlGetByID) Execute(ctx context.Context, id string) (*exampleBasicDomain.Example, error) {
	var entity postgresqlModel.ExampleModel

	err := e.GetExecutor(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exampleBasicDomain.ExampleErrNotFound
		}
		return nil, err
	}

	return exampleBasicMapper.ToExampleDomain(&entity), nil
}
