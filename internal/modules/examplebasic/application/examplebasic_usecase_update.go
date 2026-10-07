package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseUpdate struct {
	exampleUpdatePostgres  exampleBasicDomain.ExamplePostgresqlUpdate
	exampleGetByIDPostgres exampleBasicDomain.ExamplePostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseUpdate(update exampleBasicDomain.ExamplePostgresqlUpdate, getByID exampleBasicDomain.ExamplePostgresqlGetByID) exampleBasicDomain.ExampleUsecaseUpdate {
	return &ExampleUsecaseUpdate{
		exampleUpdatePostgres:  update,
		exampleGetByIDPostgres: getByID,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseUpdate) Execute(ctx context.Context, input exampleBasicDomain.Example) (*exampleBasicDomain.Example, error) {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if !existing.IsOwnedBy(input.UpdatedBy) {
		return nil, exampleBasicDomain.ExampleErrForbidden
	}

	if err := existing.UpdateExample(input.Name, input.Description, input.UpdatedBy); err != nil {
		return nil, err
	}

	if err := uc.exampleUpdatePostgres.Execute(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
