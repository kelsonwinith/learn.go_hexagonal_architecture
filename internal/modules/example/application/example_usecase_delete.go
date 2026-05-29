package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleUsecaseDelete struct {
	exampleDeletePostgres exampleDomain.ExamplePostgresqlDelete
}

func NewExampleUsecaseDelete(exampleDeletePostgres exampleDomain.ExamplePostgresqlDelete) exampleDomain.ExampleUsecaseDelete {
	return &ExampleUsecaseDelete{exampleDeletePostgres: exampleDeletePostgres}
}

func (uc *ExampleUsecaseDelete) Execute(ctx context.Context, id string) error {
	return uc.exampleDeletePostgres.Execute(ctx, id)
}
