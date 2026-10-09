package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseGetAll struct {
	exampleGetAllPostgres exampleProductDomain.ExampleProductPostgresqlGetAll
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseGetAll(exampleGetAllPostgres exampleProductDomain.ExampleProductPostgresqlGetAll) exampleProductDomain.ExampleProductUsecaseGetAll {
	return &ExampleProductUsecaseGetAll{exampleGetAllPostgres: exampleGetAllPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseGetAll) Execute(ctx context.Context) ([]*exampleProductDomain.ExampleProduct, error) {
	return uc.exampleGetAllPostgres.Execute(ctx)
}
