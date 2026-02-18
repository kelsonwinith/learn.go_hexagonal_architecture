package domain

import (
	"context"
)

type CreateExampleUseCase interface {
	Execute(ctx context.Context, input Example) (*Example, error)
}

type GetExampleByIDUseCase interface {
	Execute(ctx context.Context, id string) (*Example, error)
}

type GetAllExamplesUseCase interface {
	Execute(ctx context.Context) ([]*Example, error)
}

type UpdateExampleUseCase interface {
	Execute(ctx context.Context, input Example) (*Example, error)
}

type DeleteExampleUseCase interface {
	Execute(ctx context.Context, id string) error
}
