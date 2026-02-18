package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// Ensure GetExampleByIDUseCase implements domain.GetExampleByIDUseCase
var _ domain.GetExampleByIDUseCase = (*GetExampleByIDUseCase)(nil)

type GetExampleByIDUseCase struct {
	Repo domain.ExampleGetByIDRepository
}

func NewGetExampleByIDUseCase(repo domain.ExampleGetByIDRepository) domain.GetExampleByIDUseCase {
	return &GetExampleByIDUseCase{Repo: repo}
}

func (uc *GetExampleByIDUseCase) Execute(ctx context.Context, id string) (*domain.Example, error) {
	return uc.Repo.Execute(ctx, id)
}
