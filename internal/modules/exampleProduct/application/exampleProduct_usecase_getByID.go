package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseGetByID struct {
	exampleProductGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseGetByID(exampleProductGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID) exampleProductDomain.ExampleProductUsecaseGetByID {
	return &ExampleProductUsecaseGetByID{exampleProductGetByIDPostgres: exampleProductGetByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleProductDomain.ExampleProduct, error) {
	return uc.exampleProductGetByIDPostgres.Execute(ctx, id)
}
