package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleUsecaseCreate struct {
	exampleCreatePostgres exampleDomain.ExamplePostgresqlCreate
}

func NewExampleUsecaseCreate(exampleCreatePostgres exampleDomain.ExamplePostgresqlCreate) exampleDomain.ExampleUsecaseCreate {
	return &ExampleUsecaseCreate{exampleCreatePostgres: exampleCreatePostgres}
}

func (uc *ExampleUsecaseCreate) Execute(ctx context.Context, input exampleDomain.Example) (*exampleDomain.Example, error) {

	example, err := exampleDomain.NewExample(input.Name, input.Description)
	if err != nil {
		return nil, err
	}

	if err := uc.exampleCreatePostgres.Execute(ctx, example); err != nil {
		return nil, err
	}

	return example, nil
}
