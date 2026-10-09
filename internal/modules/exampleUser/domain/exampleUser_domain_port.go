package domain

import (
	context "context"
)

// ============================================================================
// Usecase Ports
// ============================================================================

type ExampleUserUsecaseRegister interface {
	Execute(ctx context.Context, input ExampleUser) (*ExampleUser, error)
}

type ExampleUserUsecaseLogin interface {
	Execute(ctx context.Context, email, password string) (string, error)
}

type ExampleUserUsecaseGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleUser, error)
}

// ============================================================================
// PostgreSQL Ports
// ============================================================================

type ExampleUserPostgresqlCreate interface {
	Execute(ctx context.Context, user *ExampleUser) error
}

type ExampleUserPostgresqlGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleUser, error)
}

type ExampleUserPostgresqlGetByEmail interface {
	Execute(ctx context.Context, email string) (*ExampleUser, error)
}

// ============================================================================
// Security Ports
// ============================================================================

type ExampleUserPasswordHasher interface {
	Execute(password string) (string, error)
}

type ExampleUserPasswordComparer interface {
	Execute(hashedPassword, password string) bool
}
