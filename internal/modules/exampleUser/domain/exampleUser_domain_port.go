package domain

import (
	context "context"

	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Usecase Ports
// ============================================================================

type ExampleUserUsecaseCreate interface {
	Execute(ctx context.Context, input ExampleUser) (*ExampleUser, error)
}
type ExampleUserUsecaseCreateMultiple interface {
	Execute(ctx context.Context, examples []ExampleUser) ([]*ExampleUser, error)
}
type ExampleUserUsecaseGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleUser, error)
}
type ExampleUserUsecaseGetAll interface {
	Execute(ctx context.Context) ([]*ExampleUser, error)
}
type ExampleUserUsecaseGetPaginated interface {
	Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*ExampleUser], error)
}
type ExampleUserUsecaseUpdate interface {
	Execute(ctx context.Context, input ExampleUser) (*ExampleUser, error)
}
type ExampleUserUsecaseDelete interface {
	Execute(ctx context.Context, id string, userID int64) error
}

// ============================================================================
// PostgreSQL Ports
// ============================================================================

type ExampleUserPostgresqlTransaction interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
type ExampleUserPostgresqlCreate interface {
	Execute(ctx context.Context, example *ExampleUser) error
}

type ExampleUserPostgresqlCreateMultiple interface {
	Execute(ctx context.Context, examples []*ExampleUser) error
}

type ExampleUserPostgresqlUpdate interface {
	Execute(ctx context.Context, example *ExampleUser) error
}
type ExampleUserPostgresqlDelete interface {
	Execute(ctx context.Context, id string, deletedBy int64) error
}
type ExampleUserPostgresqlGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleUser, error)
}
type ExampleUserPostgresqlGetAll interface {
	Execute(ctx context.Context) ([]*ExampleUser, error)
}
type ExampleUserPostgresqlGetPaginated interface {
	Execute(ctx context.Context, limit, offset int, search string) ([]*ExampleUser, int64, error)
}
