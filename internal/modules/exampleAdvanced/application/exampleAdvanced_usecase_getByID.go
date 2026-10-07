package application

import (
	context "context"

	exampleAdvancedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleAdvancedUsecaseGetByID struct {
	getByIDPostgres exampleAdvancedDomain.ExampleAdvancedPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleAdvancedUsecaseGetByID(getByIDPostgres exampleAdvancedDomain.ExampleAdvancedPostgresqlGetByID) exampleAdvancedDomain.ExampleAdvancedUsecaseGetByID {
	return &ExampleAdvancedUsecaseGetByID{getByIDPostgres: getByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleAdvancedUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleAdvancedDomain.Parent, error) {
	return uc.getByIDPostgres.Execute(ctx, id)
}
