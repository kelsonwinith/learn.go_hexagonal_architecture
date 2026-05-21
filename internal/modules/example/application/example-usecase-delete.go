package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type DeleteExampleUseCase struct {
	exampleDeletePostgres exampleDomain.ExampleDeletePostgres
}

func NewDeleteExampleUseCase(exampleDeletePostgres exampleDomain.ExampleDeletePostgres) exampleDomain.DeleteExampleUseCase {
	return &DeleteExampleUseCase{exampleDeletePostgres: exampleDeletePostgres}
}

func (uc *DeleteExampleUseCase) Execute(ctx context.Context, id string) error {
	return uc.exampleDeletePostgres.Execute(ctx, id)
}
