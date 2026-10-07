package domain

import (
	context "context"

	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Usecase Ports
// ============================================================================

type ExampleBasicUsecaseCreate interface {
	Execute(ctx context.Context, input ExampleBasic) (*ExampleBasic, error)
}
type ExampleBasicUsecaseCreateMultiple interface {
	Execute(ctx context.Context, examples []ExampleBasic) ([]*ExampleBasic, error)
}
type ExampleBasicUsecaseGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleBasic, error)
}
type ExampleBasicUsecaseGetAll interface {
	Execute(ctx context.Context) ([]*ExampleBasic, error)
}
type ExampleBasicUsecaseGetPaginated interface {
	Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*ExampleBasic], error)
}
type ExampleBasicUsecaseUpdate interface {
	Execute(ctx context.Context, input ExampleBasic) (*ExampleBasic, error)
}
type ExampleBasicUsecaseDelete interface {
	Execute(ctx context.Context, id string, userID int64) error
}

// ============================================================================
// PostgreSQL Ports
// ============================================================================

type ExampleBasicPostgresqlTransaction interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
type ExampleBasicPostgresqlCreate interface {
	Execute(ctx context.Context, example *ExampleBasic) error
}

type ExampleBasicPostgresqlCreateMultiple interface {
	Execute(ctx context.Context, examples []*ExampleBasic) error
}

type ExampleBasicPostgresqlUpdate interface {
	Execute(ctx context.Context, example *ExampleBasic) error
}
type ExampleBasicPostgresqlDelete interface {
	Execute(ctx context.Context, id string, deletedBy int64) error
}
type ExampleBasicPostgresqlGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleBasic, error)
}
type ExampleBasicPostgresqlGetAll interface {
	Execute(ctx context.Context) ([]*ExampleBasic, error)
}
type ExampleBasicPostgresqlGetPaginated interface {
	Execute(ctx context.Context, limit, offset int, search string) ([]*ExampleBasic, int64, error)
}
