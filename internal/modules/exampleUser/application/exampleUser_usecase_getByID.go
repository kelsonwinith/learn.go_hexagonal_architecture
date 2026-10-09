package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseGetByID struct {
	exampleUserGetByIDPostgres exampleUserDomain.ExampleUserPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseGetByID(exampleUserGetByIDPostgres exampleUserDomain.ExampleUserPostgresqlGetByID) exampleUserDomain.ExampleUserUsecaseGetByID {
	return &ExampleUserUsecaseGetByID{exampleUserGetByIDPostgres: exampleUserGetByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleUserDomain.ExampleUser, error) {
	return uc.exampleUserGetByIDPostgres.Execute(ctx, id)
}
