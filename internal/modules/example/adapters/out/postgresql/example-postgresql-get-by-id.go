package postgresql

import (
	context "context"
	errors "errors"

	postgresqlEntity "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/entity"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
	gorm "gorm.io/gorm"
)

type ExamplePostgresqlGetByID struct {
	*sharedPostgresql.Postgresql
}

func NewExamplePostgresqlGetByID(p *sharedPostgresql.Postgresql) *ExamplePostgresqlGetByID {
	return &ExamplePostgresqlGetByID{Postgresql: p}
}

func (e *ExamplePostgresqlGetByID) Execute(ctx context.Context, id string) (*exampleDomain.Example, error) {
	var entity postgresqlEntity.Example

	err := e.GetExecutor(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exampleDomain.ExampleErrNotFound
		}
		return nil, err
	}

	return toExampleDomain(&entity), nil
}
