package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type CreateExampleUseCase struct {
	Repo domain.ExampleCreateRepository
}

func NewCreateExampleUseCase(repo domain.ExampleCreateRepository) domain.CreateExampleUseCase {
	return &CreateExampleUseCase{Repo: repo}
}

func (uc *CreateExampleUseCase) Execute(ctx context.Context, input domain.Example) (*domain.Example, error) {
	example := &domain.Example{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err := uc.Repo.Execute(ctx, example)
	if err != nil {
		return nil, err
	}

	return example, nil
}
