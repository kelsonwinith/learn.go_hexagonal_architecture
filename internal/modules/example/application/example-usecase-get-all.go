package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// Ensure GetAllExamplesUseCase implements domain.GetAllExamplesUseCase
var _ domain.GetAllExamplesUseCase = (*GetAllExamplesUseCase)(nil)

type GetAllExamplesUseCase struct {
	Repo domain.ExampleGetAllRepository
}

func NewGetAllExamplesUseCase(repo domain.ExampleGetAllRepository) domain.GetAllExamplesUseCase {
	return &GetAllExamplesUseCase{Repo: repo}
}

func (uc *GetAllExamplesUseCase) Execute(ctx context.Context) ([]*domain.Example, error) {
	return uc.Repo.Execute(ctx)
}
