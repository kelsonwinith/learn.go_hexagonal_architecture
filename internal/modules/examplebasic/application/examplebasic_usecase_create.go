package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseCreate struct {
	exampleCreatePostgres exampleBasicDomain.ExamplePostgresqlCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseCreate(exampleCreatePostgres exampleBasicDomain.ExamplePostgresqlCreate) exampleBasicDomain.ExampleUsecaseCreate {
	return &ExampleUsecaseCreate{exampleCreatePostgres: exampleCreatePostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseCreate) Execute(ctx context.Context, input exampleBasicDomain.Example) (*exampleBasicDomain.Example, error) {
	example, err := exampleBasicDomain.NewExample(input.Name, input.Description, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	if err := uc.exampleCreatePostgres.Execute(ctx, example); err != nil {
		return nil, err
	}

	return example, nil
}
