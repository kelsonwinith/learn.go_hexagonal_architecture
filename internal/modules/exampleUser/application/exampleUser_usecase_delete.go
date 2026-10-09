package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseDelete struct {
	exampleDeletePostgres  exampleUserDomain.ExampleUserPostgresqlDelete
	exampleGetByIDPostgres exampleUserDomain.ExampleUserPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseDelete(
	exampleDeletePostgres exampleUserDomain.ExampleUserPostgresqlDelete,
	exampleGetByIDPostgres exampleUserDomain.ExampleUserPostgresqlGetByID,
) exampleUserDomain.ExampleUserUsecaseDelete {
	return &ExampleUserUsecaseDelete{
		exampleDeletePostgres:  exampleDeletePostgres,
		exampleGetByIDPostgres: exampleGetByIDPostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseDelete) Execute(ctx context.Context, id string, userID int64) error {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, id)
	if err != nil {
		return err
	}

	if !existing.IsOwnedBy(userID) {
		return exampleUserDomain.ExampleUserErrForbidden
	}

	return uc.exampleDeletePostgres.Execute(ctx, id, userID)
}
