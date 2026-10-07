package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseGetAll struct {
	exampleGetAllPostgres exampleBasicDomain.ExamplePostgresqlGetAll
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseGetAll(exampleGetAllPostgres exampleBasicDomain.ExamplePostgresqlGetAll) exampleBasicDomain.ExampleUsecaseGetAll {
	return &ExampleUsecaseGetAll{exampleGetAllPostgres: exampleGetAllPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseGetAll) Execute(ctx context.Context) ([]*exampleBasicDomain.Example, error) {
	return uc.exampleGetAllPostgres.Execute(ctx)
}
