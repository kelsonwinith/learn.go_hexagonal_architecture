package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicUsecaseUpdate struct {
	exampleUpdatePostgres  exampleBasicDomain.ExampleBasicPostgresqlUpdate
	exampleGetByIDPostgres exampleBasicDomain.ExampleBasicPostgresqlGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicUsecaseUpdate(update exampleBasicDomain.ExampleBasicPostgresqlUpdate, getByID exampleBasicDomain.ExampleBasicPostgresqlGetByID) exampleBasicDomain.ExampleBasicUsecaseUpdate {
	return &ExampleBasicUsecaseUpdate{
		exampleUpdatePostgres:  update,
		exampleGetByIDPostgres: getByID,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleBasicUsecaseUpdate) Execute(ctx context.Context, input exampleBasicDomain.ExampleBasic) (*exampleBasicDomain.ExampleBasic, error) {
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if !existing.IsOwnedBy(input.UpdatedBy) {
		return nil, exampleBasicDomain.ExampleBasicErrForbidden
	}

	if err := existing.UpdateExampleBasic(input.Name, input.Description, input.UpdatedBy); err != nil {
		return nil, err
	}

	if err := uc.exampleUpdatePostgres.Execute(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
