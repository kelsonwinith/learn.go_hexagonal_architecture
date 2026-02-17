package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type GetAllExamplesUseCase struct {
	Repo domain.ExampleRepository
}

func NewGetAllExamplesUseCase(repo domain.ExampleRepository) *GetAllExamplesUseCase {
	return &GetAllExamplesUseCase{Repo: repo}
}

func (uc *GetAllExamplesUseCase) Execute(ctx context.Context) ([]*domain.Example, error) {
	return uc.Repo.GetAll(ctx)
}
