package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseCreateMultiple struct {
	postgresqlTransaction  exampleBasicDomain.ExamplePostgresqlTransaction
	createMultiplePostgres exampleBasicDomain.ExamplePostgresqlCreateMultiple
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseCreateMultiple(
	postgresqlTransaction exampleBasicDomain.ExamplePostgresqlTransaction,
	createMultiplePostgres exampleBasicDomain.ExamplePostgresqlCreateMultiple,
) exampleBasicDomain.ExampleUsecaseCreateMultiple {
	return &ExampleUsecaseCreateMultiple{
		postgresqlTransaction:  postgresqlTransaction,
		createMultiplePostgres: createMultiplePostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseCreateMultiple) Execute(ctx context.Context, examples []exampleBasicDomain.Example) ([]*exampleBasicDomain.Example, error) {
	var createdExamples []*exampleBasicDomain.Example

	err := uc.postgresqlTransaction.WithinTransaction(ctx, func(ctx context.Context) error {
		createdExamples = make([]*exampleBasicDomain.Example, len(examples))
		for i := range examples {
			example, err := exampleBasicDomain.NewExample(examples[i].Name, examples[i].Description, examples[i].CreatedBy)
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
