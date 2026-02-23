package ports

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleCreate interface {
	Execute(ctx context.Context, example *domain.Example) error
}

type ExampleUpdate interface {
	Execute(ctx context.Context, example *domain.Example) error
}

type ExampleDelete interface {
	Execute(ctx context.Context, id string) error
}

type ExampleGetByID interface {
	Execute(ctx context.Context, id string) (*domain.Example, error)
}

type ExampleGetAll interface {
	Execute(ctx context.Context) ([]*domain.Example, error)
}
