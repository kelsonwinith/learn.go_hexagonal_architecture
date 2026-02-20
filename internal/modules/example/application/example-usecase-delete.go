package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/ports"
)

// Ensure DeleteExampleUseCase implements domain.DeleteExampleUseCase
var _ ports.DeleteExampleUseCase = (*DeleteExampleUseCase)(nil)

type DeleteExampleUseCase struct {
	Repo ports.ExampleDeleteRepository
}

func NewDeleteExampleUseCase(repo ports.ExampleDeleteRepository) ports.DeleteExampleUseCase {
	return &DeleteExampleUseCase{Repo: repo}
}

func (uc *DeleteExampleUseCase) Execute(ctx context.Context, id string) error {
	return uc.Repo.Execute(ctx, id)
}
