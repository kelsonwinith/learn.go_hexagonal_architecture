package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseGetByID struct {
	exampleGetByIDPostgres exampleDomain.ExamplePostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseGetByID(exampleGetByIDPostgres exampleDomain.ExamplePostgresqlGetByID) exampleDomain.ExampleUsecaseGetByID {
	return &ExampleUsecaseGetByID{exampleGetByIDPostgres: exampleGetByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleDomain.Example, error) {
	return uc.exampleGetByIDPostgres.Execute(ctx, id)
}
