package application

import (
	context "context"
	time "time"

	uuid "github.com/google/uuid"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type CreateExampleUseCase struct {
	exampleCreatePostgres exampleDomain.ExampleCreatePostgres
}

func NewCreateExampleUseCase(exampleCreatePostgres exampleDomain.ExampleCreatePostgres) exampleDomain.CreateExampleUseCase {
	return &CreateExampleUseCase{exampleCreatePostgres: exampleCreatePostgres}
}

func (uc *CreateExampleUseCase) Execute(ctx context.Context, input exampleDomain.Example) (*exampleDomain.Example, error) {
	example := &exampleDomain.Example{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err := uc.exampleCreatePostgres.Execute(ctx, example)
	if err != nil {
		return nil, err
	}

	return example, nil
}
