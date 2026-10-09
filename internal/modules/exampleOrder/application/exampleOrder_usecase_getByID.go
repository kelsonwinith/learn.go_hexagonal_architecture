package application

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderUsecaseGetByID struct {
	exampleOrderGetByIDPostgres exampleOrderDomain.ExampleOrderPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderUsecaseGetByID(exampleOrderGetByIDPostgres exampleOrderDomain.ExampleOrderPostgresqlGetByID) exampleOrderDomain.ExampleOrderUsecaseGetByID {
	return &ExampleOrderUsecaseGetByID{exampleOrderGetByIDPostgres: exampleOrderGetByIDPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleOrderUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleOrderDomain.ExampleOrder, error) {
	return uc.exampleOrderGetByIDPostgres.Execute(ctx, id)
}
