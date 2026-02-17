package application

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type UpdateExampleInput struct {
	ID          string
	Name        string
	Description string
}

type UpdateExampleUseCase struct {
	Repo domain.ExampleRepository
}

func NewUpdateExampleUseCase(repo domain.ExampleRepository) *UpdateExampleUseCase {
	return &UpdateExampleUseCase{Repo: repo}
}

func (uc *UpdateExampleUseCase) Execute(ctx context.Context, input UpdateExampleInput) (*domain.Example, error) {
	// Check if exists
	existing, err := uc.Repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	existing.Name = input.Name
	existing.Description = input.Description
	existing.UpdatedAt = time.Now().UTC()

	if err := uc.Repo.Update(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
