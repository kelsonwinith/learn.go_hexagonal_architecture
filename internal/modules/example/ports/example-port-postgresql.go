package ports

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleCreateRepository interface {
	Execute(ctx context.Context, example *domain.Example) error
}

type ExampleUpdateRepository interface {
	Execute(ctx context.Context, example *domain.Example) error
}

type ExampleDeleteRepository interface {
	Execute(ctx context.Context, id string) error
}

type ExampleGetByIDRepository interface {
	Execute(ctx context.Context, id string) (*domain.Example, error)
}

type ExampleGetAllRepository interface {
	Execute(ctx context.Context) ([]*domain.Example, error)
}
