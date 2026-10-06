package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseUpdate struct {
	exampleUpdatePostgres  exampleDomain.ExamplePostgresqlUpdate
	exampleGetByIDPostgres exampleDomain.ExamplePostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseUpdate(update exampleDomain.ExamplePostgresqlUpdate, getByID exampleDomain.ExamplePostgresqlGetByID) exampleDomain.ExampleUsecaseUpdate {
	return &ExampleUsecaseUpdate{
		exampleUpdatePostgres:  update,
		exampleGetByIDPostgres: getByID,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseUpdate) Execute(ctx context.Context, input exampleDomain.Example) (*exampleDomain.Example, error) {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if !existing.IsOwnedBy(input.UpdatedBy) {
		return nil, exampleDomain.ExampleErrForbidden
	}

	if err := existing.UpdateExample(input.Name, input.Description, input.UpdatedBy); err != nil {
		return nil, err
	}

	if err := uc.exampleUpdatePostgres.Execute(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
