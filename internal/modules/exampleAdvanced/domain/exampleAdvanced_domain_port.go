package domain

import (
	context "context"
)

// ============================================================================
// Usecase Ports
// ============================================================================

type ExampleAdvancedUsecaseCreate interface {
	Execute(ctx context.Context, input Parent) (*Parent, error)
}

type ExampleAdvancedUsecaseGetByID interface {
	Execute(ctx context.Context, id string) (*Parent, error)
}

// ============================================================================
// PostgreSQL Ports
// ============================================================================

type ExampleAdvancedPostgresqlTransaction interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type ExampleAdvancedPostgresqlCreate interface {
	Execute(ctx context.Context, parent *Parent) error
}

type ExampleAdvancedChildPostgresqlCreateMultiple interface {
	Execute(ctx context.Context, children []*Child) error
}

type ExampleAdvancedPostgresqlGetByID interface {
	Execute(ctx context.Context, id string) (*Parent, error)
}

// ============================================================================
// Event Ports
// ============================================================================

type ExampleAdvancedEventPublisher interface {
	Execute(ctx context.Context, event Event) error
}
