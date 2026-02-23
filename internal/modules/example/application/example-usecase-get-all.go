package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/ports"
)

// Ensure GetAllExamplesUseCase implements domain.GetAllExamplesUseCase
var _ ports.GetAllExamplesUseCase = (*GetAllExamplesUseCase)(nil)

type GetAllExamplesUseCase struct {
	Repo ports.ExampleGetAll
}

func NewGetAllExamplesUseCase(repo ports.ExampleGetAll) ports.GetAllExamplesUseCase {
	return &GetAllExamplesUseCase{Repo: repo}
}

func (uc *GetAllExamplesUseCase) Execute(ctx context.Context) ([]*domain.Example, error) {
	return uc.Repo.Execute(ctx)
}
