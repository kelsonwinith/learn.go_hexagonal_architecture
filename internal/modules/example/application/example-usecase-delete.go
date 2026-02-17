package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type DeleteExampleUseCase struct {
	Repo domain.ExampleRepository
}

func NewDeleteExampleUseCase(repo domain.ExampleRepository) *DeleteExampleUseCase {
	return &DeleteExampleUseCase{Repo: repo}
}

func (uc *DeleteExampleUseCase) Execute(ctx context.Context, id string) error {
	return uc.Repo.Delete(ctx, id)
}
