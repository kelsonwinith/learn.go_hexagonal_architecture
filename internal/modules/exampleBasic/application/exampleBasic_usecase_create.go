package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicUsecaseCreate struct {
	exampleCreatePostgres exampleBasicDomain.ExampleBasicPostgresqlCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicUsecaseCreate(exampleCreatePostgres exampleBasicDomain.ExampleBasicPostgresqlCreate) exampleBasicDomain.ExampleBasicUsecaseCreate {
	return &ExampleBasicUsecaseCreate{exampleCreatePostgres: exampleCreatePostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleBasicUsecaseCreate) Execute(ctx context.Context, input exampleBasicDomain.ExampleBasic) (*exampleBasicDomain.ExampleBasic, error) {
	example, err := exampleBasicDomain.NewExampleBasic(input.Name, input.Description, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	if err := uc.exampleCreatePostgres.Execute(ctx, example); err != nil {
		return nil, err
	}

	return example, nil
}
