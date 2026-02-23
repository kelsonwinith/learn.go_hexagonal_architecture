package application

import (
	context "context"

	domain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// Ensure GetAllExamplesUseCase implements domain.GetAllExamplesUseCase
var _ domain.GetAllExamplesUseCase = (*GetAllExamplesUseCase)(nil)

type GetAllExamplesUseCase struct {
	exampleGetAllPostgres domain.ExampleGetAllPostgres
}

func NewGetAllExamplesUseCase(exampleGetAllPostgres domain.ExampleGetAllPostgres) domain.GetAllExamplesUseCase {
	return &GetAllExamplesUseCase{exampleGetAllPostgres: exampleGetAllPostgres}
}

func (uc *GetAllExamplesUseCase) Execute(ctx context.Context) ([]*domain.Example, error) {
	return uc.exampleGetAllPostgres.Execute(ctx)
}
