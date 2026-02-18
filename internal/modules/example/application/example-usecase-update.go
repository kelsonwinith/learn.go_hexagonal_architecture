package application

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type UpdateExampleUseCase struct {
	Repo        domain.ExampleUpdateRepository
	GetByIDRepo domain.ExampleGetByIDRepository
}

func NewUpdateExampleUseCase(repo domain.ExampleUpdateRepository, getByIDRepo domain.ExampleGetByIDRepository) domain.UpdateExampleUseCase {
	return &UpdateExampleUseCase{
		Repo:        repo,
		GetByIDRepo: getByIDRepo,
	}
}

func (uc *UpdateExampleUseCase) Execute(ctx context.Context, input domain.Example) (*domain.Example, error) {
	// Check if exists
	existing, err := uc.GetByIDRepo.Execute(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	existing.Name = input.Name
	existing.Description = input.Description
	existing.UpdatedAt = time.Now().UTC()

	if err := uc.Repo.Execute(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
