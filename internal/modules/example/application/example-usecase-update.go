package application

import (
	context "context"
	time "time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type UpdateExampleUseCase struct {
	exampleUpdatePostgres  exampleDomain.ExampleUpdatePostgres
	exampleGetByIDPostgres exampleDomain.ExampleGetByIDPostgres
}

func NewUpdateExampleUseCase(update exampleDomain.ExampleUpdatePostgres, getByID exampleDomain.ExampleGetByIDPostgres) exampleDomain.UpdateExampleUseCase {
	return &UpdateExampleUseCase{
		exampleUpdatePostgres:  update,
		exampleGetByIDPostgres: getByID,
	}
}

func (uc *UpdateExampleUseCase) Execute(ctx context.Context, input exampleDomain.Example) (*exampleDomain.Example, error) {
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
