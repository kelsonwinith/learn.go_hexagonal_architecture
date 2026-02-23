package application

import (
	context "context"
	time "time"

	uuid "github.com/google/uuid"
	domain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type CreateExampleUseCase struct {
	exampleCreatePostgres domain.ExampleCreatePostgres
}

func NewCreateExampleUseCase(exampleCreatePostgres domain.ExampleCreatePostgres) domain.CreateExampleUseCase {
	return &CreateExampleUseCase{exampleCreatePostgres: exampleCreatePostgres}
}

func (uc *CreateExampleUseCase) Execute(ctx context.Context, input domain.Example) (*domain.Example, error) {
	example := &domain.Example{
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
