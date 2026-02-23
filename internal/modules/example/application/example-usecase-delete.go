package application

import (
	context "context"

	domain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// Ensure DeleteExampleUseCase implements domain.DeleteExampleUseCase
var _ domain.DeleteExampleUseCase = (*DeleteExampleUseCase)(nil)

type DeleteExampleUseCase struct {
	exampleDeletePostgres domain.ExampleDeletePostgres
}

func NewDeleteExampleUseCase(exampleDeletePostgres domain.ExampleDeletePostgres) domain.DeleteExampleUseCase {
	return &DeleteExampleUseCase{exampleDeletePostgres: exampleDeletePostgres}
}

func (uc *DeleteExampleUseCase) Execute(ctx context.Context, id string) error {
	return uc.exampleDeletePostgres.Execute(ctx, id)
}
