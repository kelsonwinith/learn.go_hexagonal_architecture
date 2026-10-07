package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseDelete struct {
	exampleDeletePostgres  exampleBasicDomain.ExamplePostgresqlDelete
	exampleGetByIDPostgres exampleBasicDomain.ExamplePostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseDelete(
	exampleDeletePostgres exampleBasicDomain.ExamplePostgresqlDelete,
	exampleGetByIDPostgres exampleBasicDomain.ExamplePostgresqlGetByID,
) exampleBasicDomain.ExampleUsecaseDelete {
	return &ExampleUsecaseDelete{
		exampleDeletePostgres:  exampleDeletePostgres,
		exampleGetByIDPostgres: exampleGetByIDPostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseDelete) Execute(ctx context.Context, id string, userID int64) error {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, id)
	if err != nil {
		return err
	}

	if !existing.IsOwnedBy(userID) {
		return exampleBasicDomain.ExampleErrForbidden
	}

	return uc.exampleDeletePostgres.Execute(ctx, id, userID)
}
