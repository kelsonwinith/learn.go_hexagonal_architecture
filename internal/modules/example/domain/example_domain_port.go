package domain

import (
	context "context"

	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Usecase Ports
// ============================================================================

type ExampleUsecaseCreate interface {
	Execute(ctx context.Context, input Example) (*Example, error)
}
type ExampleUsecaseCreateMultiple interface {
	Execute(ctx context.Context, examples []Example) ([]*Example, error)
}
type ExampleUsecaseGetByID interface {
	Execute(ctx context.Context, id string) (*Example, error)
}
type ExampleUsecaseGetAll interface {
	Execute(ctx context.Context) ([]*Example, error)
}
type ExampleUsecaseGetPaginated interface {
	Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*Example], error)
}
type ExampleUsecaseUpdate interface {
	Execute(ctx context.Context, input Example) (*Example, error)
}
type ExampleUsecaseDelete interface {
	Execute(ctx context.Context, id string, userID int64) error
}

// ============================================================================
// PostgreSQL Ports
// ============================================================================

type ExamplePostgresqlTransaction interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
type ExamplePostgresqlCreate interface {
	Execute(ctx context.Context, example *Example) error
}

type ExamplePostgresqlCreateMultiple interface {
	Execute(ctx context.Context, examples []*Example) error
}

type ExamplePostgresqlUpdate interface {
	Execute(ctx context.Context, example *Example) error
}
type ExamplePostgresqlDelete interface {
	Execute(ctx context.Context, id string, deletedBy int64) error
}
type ExamplePostgresqlGetByID interface {
	Execute(ctx context.Context, id string) (*Example, error)
}
type ExamplePostgresqlGetAll interface {
	Execute(ctx context.Context) ([]*Example, error)
}
type ExamplePostgresqlGetPaginated interface {
	Execute(ctx context.Context, limit, offset int, search string) ([]*Example, int64, error)
}
