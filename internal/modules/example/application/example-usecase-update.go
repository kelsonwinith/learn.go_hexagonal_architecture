package application

import (
	context "context"
	time "time"

	domain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type UpdateExampleUseCase struct {
	exampleUpdatePostgres  domain.ExampleUpdatePostgres
	exampleGetByIDPostgres domain.ExampleGetByIDPostgres
}

func NewUpdateExampleUseCase(update domain.ExampleUpdatePostgres, getByID domain.ExampleGetByIDPostgres) domain.UpdateExampleUseCase {
	return &UpdateExampleUseCase{
		exampleUpdatePostgres:  update,
		exampleGetByIDPostgres: getByID,
	}
}

func (uc *UpdateExampleUseCase) Execute(ctx context.Context, input domain.Example) (*domain.Example, error) {
	// Check if exists
	existing, err := uc.exampleGetByIDPostgres.Execute(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	existing.Name = input.Name
	existing.Description = input.Description
	existing.UpdatedAt = time.Now().UTC()

	if err := uc.exampleUpdatePostgres.Execute(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}
