package domain

import (
	context "context"
)

// Usecase Ports
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
type ExampleUsecaseUpdate interface {
	Execute(ctx context.Context, input Example) (*Example, error)
}
type ExampleUsecaseDelete interface {
	Execute(ctx context.Context, id string) error
}

// PostgreSQL Ports
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
	Execute(ctx context.Context, id string) error
}
type ExamplePostgresqlGetByID interface {
	Execute(ctx context.Context, id string) (*Example, error)
}
type ExamplePostgresqlGetAll interface {
	Execute(ctx context.Context) ([]*Example, error)
}
