package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseDelete struct {
	exampleProductDeletePostgres  exampleProductDomain.ExampleProductPostgresqlDelete
	exampleProductGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseDelete(
	exampleProductDeletePostgres exampleProductDomain.ExampleProductPostgresqlDelete,
	exampleProductGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID,
) exampleProductDomain.ExampleProductUsecaseDelete {
	return &ExampleProductUsecaseDelete{
		exampleProductDeletePostgres:  exampleProductDeletePostgres,
		exampleProductGetByIDPostgres: exampleProductGetByIDPostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseDelete) Execute(ctx context.Context, id string, userID string) error {
	exampleProduct, err := uc.exampleProductGetByIDPostgres.Execute(ctx, id)
	if err != nil {
		return err
	}

	if !exampleProduct.IsOwnedBy(userID) {
		return exampleProductDomain.ExampleProductErrForbidden
	}

	return uc.exampleProductDeletePostgres.Execute(ctx, id, userID)
}
