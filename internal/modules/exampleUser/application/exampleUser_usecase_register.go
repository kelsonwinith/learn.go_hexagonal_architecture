package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseRegister struct {
	exampleUserCreatePostgres exampleUserDomain.ExampleUserPostgresqlCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseRegister(exampleUserCreatePostgres exampleUserDomain.ExampleUserPostgresqlCreate) exampleUserDomain.ExampleUserUsecaseRegister {
	return &ExampleUserUsecaseRegister{exampleUserCreatePostgres: exampleUserCreatePostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseRegister) Execute(ctx context.Context, input exampleUserDomain.ExampleUser) (*exampleUserDomain.ExampleUser, error) {
	exampleUser, err := exampleUserDomain.NewExampleUser(input.Name, input.Email, input.Password, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	if err := uc.exampleUserCreatePostgres.Execute(ctx, exampleUser); err != nil {
		return nil, err
	}

	return exampleUser, nil
}
