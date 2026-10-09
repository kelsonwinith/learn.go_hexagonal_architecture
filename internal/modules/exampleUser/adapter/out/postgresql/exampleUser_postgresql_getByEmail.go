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

type ExampleUserPostgresqlGetByEmail struct {
	*sharedPostgresql.Postgresql
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPostgresqlGetByEmail(p *sharedPostgresql.Postgresql) *ExampleUserPostgresqlGetByEmail {
	return &ExampleUserPostgresqlGetByEmail{Postgresql: p}
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUserPostgresqlGetByEmail) Execute(ctx context.Context, email string) (*exampleUserDomain.ExampleUser, error) {
	var entity postgresqlModel.ExampleUserModel

	err := e.GetExecutor(ctx).Where("email = ?", email).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exampleUserDomain.ExampleUserErrNotFound
		}
		return nil, err
	}

	return exampleUserMapper.ToExampleUserDomain(&entity), nil
}
