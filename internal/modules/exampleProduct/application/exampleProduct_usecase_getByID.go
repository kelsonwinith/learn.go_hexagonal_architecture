package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseGetByID struct {
	exampleGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseGetByID(exampleGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID) exampleProductDomain.ExampleProductUsecaseGetByID {
	return &ExampleProductUsecaseGetByID{exampleGetByIDPostgres: exampleGetByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleProductDomain.ExampleProduct, error) {
	return uc.exampleGetByIDPostgres.Execute(ctx, id)
}
