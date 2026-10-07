package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicUsecaseCreateMultiple struct {
	postgresqlTransaction  exampleBasicDomain.ExampleBasicPostgresqlTransaction
	createMultiplePostgres exampleBasicDomain.ExampleBasicPostgresqlCreateMultiple
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicUsecaseCreateMultiple(
	postgresqlTransaction exampleBasicDomain.ExampleBasicPostgresqlTransaction,
	createMultiplePostgres exampleBasicDomain.ExampleBasicPostgresqlCreateMultiple,
) exampleBasicDomain.ExampleBasicUsecaseCreateMultiple {
	return &ExampleBasicUsecaseCreateMultiple{
		postgresqlTransaction:  postgresqlTransaction,
		createMultiplePostgres: createMultiplePostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleBasicUsecaseCreateMultiple) Execute(ctx context.Context, examples []exampleBasicDomain.ExampleBasic) ([]*exampleBasicDomain.ExampleBasic, error) {
	var createdExampleBasics []*exampleBasicDomain.ExampleBasic

	err := uc.postgresqlTransaction.WithinTransaction(ctx, func(ctx context.Context) error {
		createdExampleBasics = make([]*exampleBasicDomain.ExampleBasic, len(examples))
		for i := range examples {
			example, err := exampleBasicDomain.NewExampleBasic(examples[i].Name, examples[i].Description, examples[i].CreatedBy)
			if err != nil {
				return err
			}
			createdExampleBasics[i] = example
		}

		if err := uc.createMultiplePostgres.Execute(ctx, createdExampleBasics); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdExampleBasics, nil
}
