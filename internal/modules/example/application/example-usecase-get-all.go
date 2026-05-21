package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleUsecaseGetAll struct {
	exampleGetAllPostgres exampleDomain.ExamplePostgresqlGetAll
}

func NewExampleUsecaseGetAll(exampleGetAllPostgres exampleDomain.ExamplePostgresqlGetAll) exampleDomain.ExampleUsecaseGetAll {
	return &ExampleUsecaseGetAll{exampleGetAllPostgres: exampleGetAllPostgres}
}

func (uc *ExampleUsecaseGetAll) Execute(ctx context.Context) ([]*exampleDomain.Example, error) {
	return uc.exampleGetAllPostgres.Execute(ctx)
}
