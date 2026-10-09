package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseRegister struct {
	exampleUserPasswordHasher exampleUserDomain.ExampleUserPasswordHasher
	exampleUserCreatePostgres exampleUserDomain.ExampleUserPostgresqlCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseRegister(
	exampleUserPasswordHasher exampleUserDomain.ExampleUserPasswordHasher,
	exampleUserCreatePostgres exampleUserDomain.ExampleUserPostgresqlCreate,
) exampleUserDomain.ExampleUserUsecaseRegister {
	return &ExampleUserUsecaseRegister{
		exampleUserPasswordHasher: exampleUserPasswordHasher,
		exampleUserCreatePostgres: exampleUserCreatePostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseRegister) Execute(ctx context.Context, input exampleUserDomain.ExampleUser) (*exampleUserDomain.ExampleUser, error) {
	exampleUser, err := exampleUserDomain.NewExampleUser(input.Name, input.Email, input.Password, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := uc.exampleUserPasswordHasher.Execute(exampleUser.Password)
	if err != nil {
		return nil, err
	}
	exampleUser.Password = hashedPassword

	if err := uc.exampleUserCreatePostgres.Execute(ctx, exampleUser); err != nil {
		return nil, err
	}

	return exampleUser, nil
}
