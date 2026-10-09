package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseCreate struct {
	exampleProductCreatePostgres exampleProductDomain.ExampleProductPostgresqlCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseCreate(exampleProductCreatePostgres exampleProductDomain.ExampleProductPostgresqlCreate) exampleProductDomain.ExampleProductUsecaseCreate {
	return &ExampleProductUsecaseCreate{exampleProductCreatePostgres: exampleProductCreatePostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseCreate) Execute(ctx context.Context, exampleProductInput exampleProductDomain.ExampleProduct) (*exampleProductDomain.ExampleProduct, error) {
	exampleProduct, err := exampleProductDomain.NewExampleProduct(exampleProductInput.Name, exampleProductInput.Description, exampleProductInput.Price, exampleProductInput.CreatedBy)
	if err != nil {
		return nil, err
	}

	if err := uc.exampleProductCreatePostgres.Execute(ctx, exampleProduct); err != nil {
		return nil, err
	}

	return exampleProduct, nil
}
