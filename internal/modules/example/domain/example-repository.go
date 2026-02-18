package domain

import "context"

type ExampleCreateRepository interface {
	Execute(ctx context.Context, example *Example) error
}

type ExampleUpdateRepository interface {
	Execute(ctx context.Context, example *Example) error
}

type ExampleDeleteRepository interface {
	Execute(ctx context.Context, id string) error
}

type ExampleGetByIDRepository interface {
	Execute(ctx context.Context, id string) (*Example, error)
}

type ExampleGetAllRepository interface {
	Execute(ctx context.Context) ([]*Example, error)
}
