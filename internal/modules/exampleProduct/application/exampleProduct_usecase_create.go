package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseCreate struct {
	exampleCreatePostgres exampleProductDomain.ExampleProductPostgresqlCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseCreate(exampleCreatePostgres exampleProductDomain.ExampleProductPostgresqlCreate) exampleProductDomain.ExampleProductUsecaseCreate {
	return &ExampleProductUsecaseCreate{exampleCreatePostgres: exampleCreatePostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseCreate) Execute(ctx context.Context, input exampleProductDomain.ExampleProduct) (*exampleProductDomain.ExampleProduct, error) {
	example, err := exampleProductDomain.NewExampleProduct(input.Name, input.Description, input.Price, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	if err := uc.exampleCreatePostgres.Execute(ctx, example); err != nil {
		return nil, err
	}

	return example, nil
}
