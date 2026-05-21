package application

import (
	context "context"
	time "time"

	uuid "github.com/google/uuid"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

type ExampleUsecaseCreateMultiple struct {
	createMultiplePostgres exampleDomain.ExamplePostgresqlCreateMultiple
	txManager              sharedDomain.TransactionManagerPostgresql
}

func NewExampleUsecaseCreateMultiple(
	createMultiplePostgres exampleDomain.ExamplePostgresqlCreateMultiple,
	txManager sharedDomain.TransactionManagerPostgresql,
) exampleDomain.ExampleUsecaseCreateMultiple {
	return &ExampleUsecaseCreateMultiple{
		createMultiplePostgres: createMultiplePostgres,
		txManager:              txManager,
	}
}

func (uc *ExampleUsecaseCreateMultiple) Execute(ctx context.Context, examples []exampleDomain.Example) ([]*exampleDomain.Example, error) {
	var createdExamples []*exampleDomain.Example

	err := uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {
		now := time.Now().UTC()
		createdExamples = make([]*exampleDomain.Example, len(examples))
		for i := range examples {
			createdExamples[i] = &exampleDomain.Example{
				ID:          uuid.New().String(),
				Name:        examples[i].Name,
				Description: examples[i].Description,
				CreatedAt:   now,
				UpdatedAt:   now,
			}
		}

		if err := uc.createMultiplePostgres.Execute(ctx, createdExamples); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdExamples, nil
}
