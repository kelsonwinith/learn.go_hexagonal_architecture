package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// Ensure DeleteExampleUseCase implements domain.DeleteExampleUseCase
var _ domain.DeleteExampleUseCase = (*DeleteExampleUseCase)(nil)

type DeleteExampleUseCase struct {
	Repo domain.ExampleDeleteRepository
}

func NewDeleteExampleUseCase(repo domain.ExampleDeleteRepository) domain.DeleteExampleUseCase {
	return &DeleteExampleUseCase{Repo: repo}
}

func (uc *DeleteExampleUseCase) Execute(ctx context.Context, id string) error {
	return uc.Repo.Execute(ctx, id)
}
