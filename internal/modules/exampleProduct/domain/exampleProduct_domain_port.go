package domain

import (
	context "context"

	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Usecase Ports
// ============================================================================

type ExampleProductUsecaseCreate interface {
	Execute(ctx context.Context, input ExampleProduct) (*ExampleProduct, error)
}

type ExampleProductUsecaseGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleProduct, error)
}

type ExampleProductUsecaseGetPaginated interface {
	Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*ExampleProduct], error)
}

type ExampleProductUsecaseUpdate interface {
	Execute(ctx context.Context, input ExampleProduct) (*ExampleProduct, error)
}

type ExampleProductUsecaseDelete interface {
	Execute(ctx context.Context, id string, userID int64) error
}

// ============================================================================
// PostgreSQL Ports
// ============================================================================

type ExampleProductPostgresqlCreate interface {
	Execute(ctx context.Context, product *ExampleProduct) error
}

type ExampleProductPostgresqlUpdate interface {
	Execute(ctx context.Context, product *ExampleProduct) error
}

type ExampleProductPostgresqlDelete interface {
	Execute(ctx context.Context, id string, deletedBy int64) error
}

type ExampleProductPostgresqlGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleProduct, error)
}

type ExampleProductPostgresqlGetPaginated interface {
	Execute(ctx context.Context, limit, offset int, search string) ([]*ExampleProduct, int64, error)
}
