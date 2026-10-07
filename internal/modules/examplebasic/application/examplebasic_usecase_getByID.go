package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseGetByID struct {
	exampleGetByIDPostgres exampleBasicDomain.ExamplePostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseGetByID(exampleGetByIDPostgres exampleBasicDomain.ExamplePostgresqlGetByID) exampleBasicDomain.ExampleUsecaseGetByID {
	return &ExampleUsecaseGetByID{exampleGetByIDPostgres: exampleGetByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleBasicDomain.Example, error) {
	return uc.exampleGetByIDPostgres.Execute(ctx, id)
}
