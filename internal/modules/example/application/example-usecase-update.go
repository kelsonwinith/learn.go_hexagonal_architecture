package application

import (
	"context"
	"time"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/ports"
)

type UpdateExampleUseCase struct {
	UpdateRepo  ports.ExampleUpdateRepository
	GetByIDRepo ports.ExampleGetByIDRepository
}

func NewUpdateExampleUseCase(updateRepo ports.ExampleUpdateRepository, getByIDRepo ports.ExampleGetByIDRepository) ports.UpdateExampleUseCase {
	return &UpdateExampleUseCase{
		UpdateRepo:  updateRepo,
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

	if err := uc.UpdateRepo.Execute(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
