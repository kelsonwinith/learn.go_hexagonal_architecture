package ports

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type CreateExampleUseCase interface {
	Execute(ctx context.Context, input domain.Example) (*domain.Example, error)
}

type GetExampleByIDUseCase interface {
	Execute(ctx context.Context, id string) (*domain.Example, error)
}

type GetAllExamplesUseCase interface {
	Execute(ctx context.Context) ([]*domain.Example, error)
}

type UpdateExampleUseCase interface {
	Execute(ctx context.Context, input domain.Example) (*domain.Example, error)
}

type DeleteExampleUseCase interface {
	Execute(ctx context.Context, id string) error
}
