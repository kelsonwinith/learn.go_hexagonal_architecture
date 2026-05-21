package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type GetExampleByIDUseCase struct {
	exampleGetByIDPostgres exampleDomain.ExampleGetByIDPostgres
}

func NewGetExampleByIDUseCase(exampleGetByIDPostgres exampleDomain.ExampleGetByIDPostgres) exampleDomain.GetExampleByIDUseCase {
	return &GetExampleByIDUseCase{exampleGetByIDPostgres: exampleGetByIDPostgres}
}

func (uc *GetExampleByIDUseCase) Execute(ctx context.Context, id string) (*exampleDomain.Example, error) {
	return uc.exampleGetByIDPostgres.Execute(ctx, id)
}
