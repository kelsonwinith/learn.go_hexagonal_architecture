package application

import (
	context "context"
	time "time"

	uuid "github.com/google/uuid"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleUsecaseCreate struct {
	exampleCreatePostgres exampleDomain.ExamplePostgresqlCreate
}

func NewExampleUsecaseCreate(exampleCreatePostgres exampleDomain.ExamplePostgresqlCreate) exampleDomain.ExampleUsecaseCreate {
	return &ExampleUsecaseCreate{exampleCreatePostgres: exampleCreatePostgres}
}

func (uc *ExampleUsecaseCreate) Execute(ctx context.Context, input exampleDomain.Example) (*exampleDomain.Example, error) {
	example := &exampleDomain.Example{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err := uc.exampleCreatePostgres.Execute(ctx, example)
	if err != nil {
		return nil, err
	}

	return example, nil
}
