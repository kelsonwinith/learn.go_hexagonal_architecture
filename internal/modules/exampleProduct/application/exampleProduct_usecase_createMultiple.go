package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseCreateMultiple struct {
	postgresqlTransaction  exampleProductDomain.ExampleProductPostgresqlTransaction
	createMultiplePostgres exampleProductDomain.ExampleProductPostgresqlCreateMultiple
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseCreateMultiple(
	postgresqlTransaction exampleProductDomain.ExampleProductPostgresqlTransaction,
	createMultiplePostgres exampleProductDomain.ExampleProductPostgresqlCreateMultiple,
) exampleProductDomain.ExampleProductUsecaseCreateMultiple {
	return &ExampleProductUsecaseCreateMultiple{
		postgresqlTransaction:  postgresqlTransaction,
		createMultiplePostgres: createMultiplePostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseCreateMultiple) Execute(ctx context.Context, examples []exampleProductDomain.ExampleProduct) ([]*exampleProductDomain.ExampleProduct, error) {
	var createdExampleProducts []*exampleProductDomain.ExampleProduct

	err := uc.postgresqlTransaction.WithinTransaction(ctx, func(ctx context.Context) error {
		createdExampleProducts = make([]*exampleProductDomain.ExampleProduct, len(examples))
		for i := range examples {
			example, err := exampleProductDomain.NewExampleProduct(examples[i].Name, examples[i].Description, examples[i].Price, examples[i].CreatedBy)
			if err != nil {
				return err
			}
			createdExampleProducts[i] = example
		}

		if err := uc.createMultiplePostgres.Execute(ctx, createdExampleProducts); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdExampleProducts, nil
}
