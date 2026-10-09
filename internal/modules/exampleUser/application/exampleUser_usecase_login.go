package application

import (
	context "context"
	strings "strings"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseLogin struct {
	exampleUserPasswordComparer   exampleUserDomain.ExampleUserPasswordComparer
	exampleUserGetByEmailPostgres exampleUserDomain.ExampleUserPostgresqlGetByEmail
	exampleUserTokenService       sharedDomain.TokenService
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseLogin(
	exampleUserPasswordComparer exampleUserDomain.ExampleUserPasswordComparer,
	exampleUserGetByEmailPostgres exampleUserDomain.ExampleUserPostgresqlGetByEmail,
	exampleUserTokenService sharedDomain.TokenService,
) exampleUserDomain.ExampleUserUsecaseLogin {
	return &ExampleUserUsecaseLogin{
		exampleUserPasswordComparer:   exampleUserPasswordComparer,
		exampleUserGetByEmailPostgres: exampleUserGetByEmailPostgres,
		exampleUserTokenService:       exampleUserTokenService,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseLogin) Execute(ctx context.Context, email, password string) (string, error) {
	exampleUser, err := uc.exampleUserGetByEmailPostgres.Execute(ctx, strings.ToLower(strings.TrimSpace(email)))
	if err != nil {
		return "", exampleUserDomain.ExampleUserErrInvalidCredentials
	}

	if !uc.exampleUserPasswordComparer.Execute(exampleUser.Password, password) {
		return "", exampleUserDomain.ExampleUserErrInvalidCredentials
	}

	token, err := uc.exampleUserTokenService.Generate(ctx, exampleUser.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
