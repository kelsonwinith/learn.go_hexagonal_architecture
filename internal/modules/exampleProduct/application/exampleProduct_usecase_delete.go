package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseDelete struct {
	exampleDeletePostgres  exampleProductDomain.ExampleProductPostgresqlDelete
	exampleGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseDelete(
	exampleDeletePostgres exampleProductDomain.ExampleProductPostgresqlDelete,
	exampleGetByIDPostgres exampleProductDomain.ExampleProductPostgresqlGetByID,
) exampleProductDomain.ExampleProductUsecaseDelete {
	return &ExampleProductUsecaseDelete{
		exampleDeletePostgres:  exampleDeletePostgres,
		exampleGetByIDPostgres: exampleGetByIDPostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseDelete) Execute(ctx context.Context, id string, userID int64) error {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, id)
	if err != nil {
		return err
	}

	if !existing.IsOwnedBy(userID) {
		return exampleProductDomain.ExampleProductErrForbidden
	}

	return uc.exampleDeletePostgres.Execute(ctx, id, userID)
}
