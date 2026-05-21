package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type GetAllExamplesUseCase struct {
	exampleGetAllPostgres exampleDomain.ExampleGetAllPostgres
}

func NewGetAllExamplesUseCase(exampleGetAllPostgres exampleDomain.ExampleGetAllPostgres) exampleDomain.GetAllExamplesUseCase {
	return &GetAllExamplesUseCase{exampleGetAllPostgres: exampleGetAllPostgres}
}

func (uc *GetAllExamplesUseCase) Execute(ctx context.Context) ([]*exampleDomain.Example, error) {
	return uc.exampleGetAllPostgres.Execute(ctx)
}
