package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseUpdate struct {
	exampleUpdatePostgres  exampleProductDomain.ExampleProductPostgresqlUpdate
	exampleGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseUpdate(update exampleProductDomain.ExampleProductPostgresqlUpdate, getByID exampleProductDomain.ExampleProductPostgresqlGetByID) exampleProductDomain.ExampleProductUsecaseUpdate {
	return &ExampleProductUsecaseUpdate{
		exampleUpdatePostgres:  update,
		exampleGetByIDPostgres: getByID,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseUpdate) Execute(ctx context.Context, input exampleProductDomain.ExampleProduct) (*exampleProductDomain.ExampleProduct, error) {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if !existing.IsOwnedBy(input.UpdatedBy) {
		return nil, exampleProductDomain.ExampleProductErrForbidden
	}

	if err := existing.UpdateExampleProduct(input.Name, input.Description, input.Price, input.UpdatedBy); err != nil {
		return nil, err
	}

	if err := uc.exampleUpdatePostgres.Execute(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
