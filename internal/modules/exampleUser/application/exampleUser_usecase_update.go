package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseUpdate struct {
	exampleUpdatePostgres  exampleUserDomain.ExampleUserPostgresqlUpdate
	exampleGetByIDPostgres exampleUserDomain.ExampleUserPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseUpdate(update exampleUserDomain.ExampleUserPostgresqlUpdate, getByID exampleUserDomain.ExampleUserPostgresqlGetByID) exampleUserDomain.ExampleUserUsecaseUpdate {
	return &ExampleUserUsecaseUpdate{
		exampleUpdatePostgres:  update,
		exampleGetByIDPostgres: getByID,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseUpdate) Execute(ctx context.Context, input exampleUserDomain.ExampleUser) (*exampleUserDomain.ExampleUser, error) {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if !existing.IsOwnedBy(input.UpdatedBy) {
		return nil, exampleUserDomain.ExampleUserErrForbidden
	}

	if err := existing.UpdateExampleUser(input.Name, input.Description, input.UpdatedBy); err != nil {
		return nil, err
	}

	if err := uc.exampleUpdatePostgres.Execute(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
