package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseUpdate struct {
	exampleProductUpdatePostgres  exampleProductDomain.ExampleProductPostgresqlUpdate
	exampleProductGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseUpdate(exampleProductUpdatePostgres exampleProductDomain.ExampleProductPostgresqlUpdate, exampleProductGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID) exampleProductDomain.ExampleProductUsecaseUpdate {
	return &ExampleProductUsecaseUpdate{
		exampleProductUpdatePostgres:  exampleProductUpdatePostgres,
		exampleProductGetByIDPostgres: exampleProductGetByIDPostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseUpdate) Execute(ctx context.Context, exampleProductInput exampleProductDomain.ExampleProduct) (*exampleProductDomain.ExampleProduct, error) {
	exampleProduct, err := uc.exampleProductGetByIDPostgres.Execute(ctx, exampleProductInput.ID)
	if err != nil {
		return nil, err
	}

	if !exampleProduct.IsOwnedBy(exampleProductInput.UpdatedBy) {
		return nil, exampleProductDomain.ExampleProductErrForbidden
	}

	if err := exampleProduct.UpdateExampleProduct(exampleProductInput.Name, exampleProductInput.Description, exampleProductInput.Price, exampleProductInput.UpdatedBy); err != nil {
		return nil, err
	}

	if err := uc.exampleProductUpdatePostgres.Execute(ctx, exampleProduct); err != nil {
		return nil, err
	}

	return exampleProduct, nil
}
