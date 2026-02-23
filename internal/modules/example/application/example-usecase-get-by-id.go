package application

import (
	"context"

	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/ports"
)

// Ensure GetExampleByIDUseCase implements domain.GetExampleByIDUseCase
var _ ports.GetExampleByIDUseCase = (*GetExampleByIDUseCase)(nil)

type GetExampleByIDUseCase struct {
	Repo ports.ExampleGetByID
}

func NewGetExampleByIDUseCase(repo ports.ExampleGetByID) ports.GetExampleByIDUseCase {
	return &GetExampleByIDUseCase{Repo: repo}
}

func (uc *GetExampleByIDUseCase) Execute(ctx context.Context, id string) (*domain.Example, error) {
	return uc.Repo.Execute(ctx, id)
}
