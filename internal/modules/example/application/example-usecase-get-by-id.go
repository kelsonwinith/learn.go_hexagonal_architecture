package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type GetExampleByIDUseCase struct {
	Repo domain.ExampleRepository
}

func NewGetExampleByIDUseCase(repo domain.ExampleRepository) *GetExampleByIDUseCase {
	return &GetExampleByIDUseCase{Repo: repo}
}

func (uc *GetExampleByIDUseCase) Execute(ctx context.Context, id string) (*domain.Example, error) {
	return uc.Repo.GetByID(ctx, id)
}
