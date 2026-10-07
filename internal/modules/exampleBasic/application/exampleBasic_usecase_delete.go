package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicUsecaseDelete struct {
	exampleDeletePostgres  exampleBasicDomain.ExampleBasicPostgresqlDelete
	exampleGetByIDPostgres exampleBasicDomain.ExampleBasicPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicUsecaseDelete(
	exampleDeletePostgres exampleBasicDomain.ExampleBasicPostgresqlDelete,
	exampleGetByIDPostgres exampleBasicDomain.ExampleBasicPostgresqlGetByID,
) exampleBasicDomain.ExampleBasicUsecaseDelete {
	return &ExampleBasicUsecaseDelete{
		exampleDeletePostgres:  exampleDeletePostgres,
		exampleGetByIDPostgres: exampleGetByIDPostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleBasicUsecaseDelete) Execute(ctx context.Context, id string, userID int64) error {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, id)
	if err != nil {
		return err
	}

	if !existing.IsOwnedBy(userID) {
		return exampleBasicDomain.ExampleBasicErrForbidden
	}

	return uc.exampleDeletePostgres.Execute(ctx, id, userID)
}
