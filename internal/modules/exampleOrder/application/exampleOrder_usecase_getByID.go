package application

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderUsecaseGetByID struct {
	getByIDPostgres exampleOrderDomain.ExampleOrderPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderUsecaseGetByID(getByIDPostgres exampleOrderDomain.ExampleOrderPostgresqlGetByID) exampleOrderDomain.ExampleOrderUsecaseGetByID {
	return &ExampleOrderUsecaseGetByID{getByIDPostgres: getByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleOrderUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleOrderDomain.ExampleOrder, error) {
	return uc.getByIDPostgres.Execute(ctx, id)
}
