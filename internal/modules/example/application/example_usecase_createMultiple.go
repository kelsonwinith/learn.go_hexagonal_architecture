package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
)

type ExampleUsecaseCreateMultiple struct {
	postgresqlTransaction  *sharedPostgresql.PostgresqlTransaction
	createMultiplePostgres exampleDomain.ExamplePostgresqlCreateMultiple
}

func NewExampleUsecaseCreateMultiple(
	postgresqlTransaction *sharedPostgresql.PostgresqlTransaction,
	createMultiplePostgres exampleDomain.ExamplePostgresqlCreateMultiple,
) exampleDomain.ExampleUsecaseCreateMultiple {
	return &ExampleUsecaseCreateMultiple{
		postgresqlTransaction:  postgresqlTransaction,
		createMultiplePostgres: createMultiplePostgres,
	}
}

func (uc *ExampleUsecaseCreateMultiple) Execute(ctx context.Context, examples []exampleDomain.Example) ([]*exampleDomain.Example, error) {
	var createdExamples []*exampleDomain.Example

	err := uc.postgresqlTransaction.WithinTransaction(ctx, func(ctx context.Context) error {
		createdExamples = make([]*exampleDomain.Example, len(examples))
		for i := range examples {
			example, err := exampleDomain.NewExample(examples[i].Name, examples[i].Description)
			if err != nil {
				return err
			}
			createdExamples[i] = example
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
