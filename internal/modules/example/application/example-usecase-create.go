package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type CreateExampleInput struct {
	Name        string
	Description string
}

type CreateExampleUseCase struct {
	Repo domain.ExampleRepository
}

func NewCreateExampleUseCase(repo domain.ExampleRepository) *CreateExampleUseCase {
	return &CreateExampleUseCase{Repo: repo}
}

func (uc *CreateExampleUseCase) Execute(ctx context.Context, input CreateExampleInput) (*domain.Example, error) {
	example := &domain.Example{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := uc.Repo.Create(ctx, example); err != nil {
		return nil, err
	}

	return example, nil
}
