package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicUsecaseGetByID struct {
	exampleGetByIDPostgres exampleBasicDomain.ExampleBasicPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicUsecaseGetByID(exampleGetByIDPostgres exampleBasicDomain.ExampleBasicPostgresqlGetByID) exampleBasicDomain.ExampleBasicUsecaseGetByID {
	return &ExampleBasicUsecaseGetByID{exampleGetByIDPostgres: exampleGetByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleBasicUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleBasicDomain.ExampleBasic, error) {
	return uc.exampleGetByIDPostgres.Execute(ctx, id)
}
