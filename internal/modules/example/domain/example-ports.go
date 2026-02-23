package domain

import (
	"context"
)

// Usecase Ports
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

// PostgreSQL Ports
type ExampleCreatePostgres interface {
	Execute(ctx context.Context, example *Example) error
}
type ExampleUpdatePostgres interface {
	Execute(ctx context.Context, example *Example) error
}
type ExampleDeletePostgres interface {
	Execute(ctx context.Context, id string) error
}
type ExampleGetByIDPostgres interface {
	Execute(ctx context.Context, id string) (*Example, error)
}
type ExampleGetAllPostgres interface {
	Execute(ctx context.Context) ([]*Example, error)
}
