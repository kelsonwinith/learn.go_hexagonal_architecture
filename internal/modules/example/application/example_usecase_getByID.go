package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleUsecaseGetByID struct {
	exampleGetByIDPostgres exampleDomain.ExamplePostgresqlGetByID
}

func NewExampleUsecaseGetByID(exampleGetByIDPostgres exampleDomain.ExamplePostgresqlGetByID) exampleDomain.ExampleUsecaseGetByID {
	return &ExampleUsecaseGetByID{exampleGetByIDPostgres: exampleGetByIDPostgres}
}

func (uc *ExampleUsecaseGetByID) Execute(ctx context.Context, id string) (*exampleDomain.Example, error) {
	return uc.exampleGetByIDPostgres.Execute(ctx, id)
}
