package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseCreate struct {
	exampleCreatePostgres exampleUserDomain.ExampleUserPostgresqlCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseCreate(exampleCreatePostgres exampleUserDomain.ExampleUserPostgresqlCreate) exampleUserDomain.ExampleUserUsecaseCreate {
	return &ExampleUserUsecaseCreate{exampleCreatePostgres: exampleCreatePostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseCreate) Execute(ctx context.Context, input exampleUserDomain.ExampleUser) (*exampleUserDomain.ExampleUser, error) {
	example, err := exampleUserDomain.NewExampleUser(input.Name, input.Description, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	if err := uc.exampleCreatePostgres.Execute(ctx, example); err != nil {
		return nil, err
	}

	return example, nil
}
