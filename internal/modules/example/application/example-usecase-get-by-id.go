package application

import (
	context "context"

	domain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// Ensure GetExampleByIDUseCase implements domain.GetExampleByIDUseCase
var _ domain.GetExampleByIDUseCase = (*GetExampleByIDUseCase)(nil)

type GetExampleByIDUseCase struct {
	exampleGetByIDPostgres domain.ExampleGetByIDPostgres
}

func NewGetExampleByIDUseCase(exampleGetByIDPostgres domain.ExampleGetByIDPostgres) domain.GetExampleByIDUseCase {
	return &GetExampleByIDUseCase{exampleGetByIDPostgres: exampleGetByIDPostgres}
}

func (uc *GetExampleByIDUseCase) Execute(ctx context.Context, id string) (*domain.Example, error) {
	return uc.exampleGetByIDPostgres.Execute(ctx, id)
}
