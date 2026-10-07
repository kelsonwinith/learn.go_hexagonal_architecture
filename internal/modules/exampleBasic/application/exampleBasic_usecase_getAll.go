package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicUsecaseGetAll struct {
	exampleGetAllPostgres exampleBasicDomain.ExampleBasicPostgresqlGetAll
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicUsecaseGetAll(exampleGetAllPostgres exampleBasicDomain.ExampleBasicPostgresqlGetAll) exampleBasicDomain.ExampleBasicUsecaseGetAll {
	return &ExampleBasicUsecaseGetAll{exampleGetAllPostgres: exampleGetAllPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleBasicUsecaseGetAll) Execute(ctx context.Context) ([]*exampleBasicDomain.ExampleBasic, error) {
	return uc.exampleGetAllPostgres.Execute(ctx)
}
