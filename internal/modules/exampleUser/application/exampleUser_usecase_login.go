package application

import (
	context "context"
	strings "strings"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseLogin struct {
	exampleUserGetByEmailPostgres exampleUserDomain.ExampleUserPostgresqlGetByEmail
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseLogin(exampleUserGetByEmailPostgres exampleUserDomain.ExampleUserPostgresqlGetByEmail) exampleUserDomain.ExampleUserUsecaseLogin {
	return &ExampleUserUsecaseLogin{exampleUserGetByEmailPostgres: exampleUserGetByEmailPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseLogin) Execute(ctx context.Context, email, password string) (*exampleUserDomain.ExampleUser, error) {
	exampleUser, err := uc.exampleUserGetByEmailPostgres.Execute(ctx, strings.ToLower(strings.TrimSpace(email)))
	if err != nil {
		return nil, exampleUserDomain.ExampleUserErrInvalidCredentials
	}

	// Mock authentication: plain comparison, no hashing.
	if exampleUser.Password != password {
		return nil, exampleUserDomain.ExampleUserErrInvalidCredentials
	}

	return exampleUser, nil
}
