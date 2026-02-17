package domain

import "context"

type ExampleRepository interface {
	Create(ctx context.Context, example *Example) error
	GetAll(ctx context.Context) ([]*Example, error)
	GetByID(ctx context.Context, id string) (*Example, error)
	Update(ctx context.Context, example *Example) error
	Delete(ctx context.Context, id string) error
}
