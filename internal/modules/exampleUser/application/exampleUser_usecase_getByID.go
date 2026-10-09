package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseGetByID struct {
	exampleGetByIDPostgres exampleUserDomain.ExampleUserPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseGetByID(exampleGetByIDPostgres exampleUserDomain.ExampleUserPostgresqlGetByID) exampleUserDomain.ExampleUserUsecaseGetByID {
	return &ExampleUserUsecaseGetByID{exampleGetByIDPostgres: exampleGetByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleUserDomain.ExampleUser, error) {
	return uc.exampleGetByIDPostgres.Execute(ctx, id)
}
