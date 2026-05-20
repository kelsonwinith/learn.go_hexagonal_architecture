package application

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	domain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

type CreateMultipleExamplesUseCase struct {
	createMultiplePostgres domain.ExampleCreateMultiplePostgres
	postgres               *postgresql.Postgresql
}

func NewCreateMultipleExamplesUseCase(createMultiplePostgres domain.ExampleCreateMultiplePostgres, postgres *postgresql.Postgresql) domain.CreateMultipleExamplesUseCase {
	return &CreateMultipleExamplesUseCase{
		createMultiplePostgres: createMultiplePostgres,
		postgres:               postgres,
	}
}

func (uc *CreateMultipleExamplesUseCase) Execute(ctx context.Context, examples []domain.Example) ([]*domain.Example, error) {
	var createdExamples []*domain.Example

	err := uc.postgres.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		now := time.Now().UTC()
		for i := range examples {
			examples[i].CreatedAt = now
			examples[i].UpdatedAt = now
		}

		createdExamples = make([]*domain.Example, len(examples))
		if err := uc.createMultiplePostgres.Execute(ctx, tx, createdExamples); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdExamples, nil
}
