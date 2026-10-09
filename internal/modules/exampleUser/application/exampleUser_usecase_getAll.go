package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseGetAll struct {
	exampleGetAllPostgres exampleUserDomain.ExampleUserPostgresqlGetAll
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseGetAll(exampleGetAllPostgres exampleUserDomain.ExampleUserPostgresqlGetAll) exampleUserDomain.ExampleUserUsecaseGetAll {
	return &ExampleUserUsecaseGetAll{exampleGetAllPostgres: exampleGetAllPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseGetAll) Execute(ctx context.Context) ([]*exampleUserDomain.ExampleUser, error) {
	return uc.exampleGetAllPostgres.Execute(ctx)
}
