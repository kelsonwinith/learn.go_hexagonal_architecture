package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseDelete struct {
	exampleDeletePostgres  exampleDomain.ExamplePostgresqlDelete
	exampleGetByIDPostgres exampleDomain.ExamplePostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseDelete(
	exampleDeletePostgres exampleDomain.ExamplePostgresqlDelete,
	exampleGetByIDPostgres exampleDomain.ExamplePostgresqlGetByID,
) exampleDomain.ExampleUsecaseDelete {
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

	if existing.CreatedBy != userID {
		return exampleDomain.ExampleErrForbidden
	}

	return uc.exampleDeletePostgres.Execute(ctx, id, userID)
}
