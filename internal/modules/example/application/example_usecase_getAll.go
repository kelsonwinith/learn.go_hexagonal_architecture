package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseGetAll struct {
	exampleGetAllPostgres exampleDomain.ExamplePostgresqlGetAll
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseGetAll(exampleGetAllPostgres exampleDomain.ExamplePostgresqlGetAll) exampleDomain.ExampleUsecaseGetAll {
	return &ExampleUsecaseGetAll{exampleGetAllPostgres: exampleGetAllPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseGetAll) Execute(ctx context.Context) ([]*exampleDomain.Example, error) {
	return uc.exampleGetAllPostgres.Execute(ctx)
}
