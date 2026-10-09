package domain

import (
	context "context"

	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Usecase Ports
// ============================================================================

type ExampleOrderUsecaseCreate interface {
	Execute(ctx context.Context, input ExampleOrder) (*ExampleOrder, error)
}

type ExampleOrderUsecaseGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleOrder, error)
}

type ExampleOrderUsecaseGetPaginated interface {
	Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*ExampleOrder], error)
}

// ============================================================================
// PostgreSQL Ports
// ============================================================================

type ExampleOrderPostgresqlTransaction interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type ExampleOrderPostgresqlCreate interface {
	Execute(ctx context.Context, order *ExampleOrder) error
}

type ExampleOrderProductPostgresqlCreateMultiple interface {
	Execute(ctx context.Context, products []*ExampleOrderProduct) error
}

type ExampleOrderPostgresqlGetByID interface {
	Execute(ctx context.Context, id string) (*ExampleOrder, error)
}

type ExampleOrderPostgresqlGetPaginated interface {
	Execute(ctx context.Context, limit, offset int, search string) ([]*ExampleOrder, int64, error)
}

// ============================================================================
// Event Ports
// ============================================================================

type ExampleOrderEventPublisher interface {
	Execute(ctx context.Context, event Event) error
}

// ============================================================================
// Cross-module Ports
// ============================================================================

type ExampleOrderUser struct {
	ID   string
	Name string
}

type ExampleOrderProductInfo struct {
	ID    string
	Name  string
	Price int64
}

type ExampleUserModuleGetByID interface {
	Execute(ctx context.Context, id string) (ExampleOrderUser, error)
}

type ExampleProductModuleGetByID interface {
	Execute(ctx context.Context, id string) (ExampleOrderProductInfo, error)
}
