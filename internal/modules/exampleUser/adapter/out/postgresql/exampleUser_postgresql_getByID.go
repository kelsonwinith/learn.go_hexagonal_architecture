package postgresql

import (
	context "context"
	errors "errors"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	exampleUserMapper "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/out/postgresql/mapper"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserPostgresqlGetByID struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPostgresqlGetByID(p *sharedPostgresql.Postgresql) *ExampleUserPostgresqlGetByID {
	return &ExampleUserPostgresqlGetByID{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUserPostgresqlGetByID) Execute(ctx context.Context, id string) (*exampleUserDomain.ExampleUser, error) {
	var entity postgresqlModel.ExampleUserModel

	err := e.GetExecutor(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exampleUserDomain.ExampleUserErrNotFound
		}
		return nil, err
	}

	return exampleUserMapper.ToExampleUserDomain(&entity), nil
}
