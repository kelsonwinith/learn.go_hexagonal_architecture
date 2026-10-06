package application

import (
	context "context"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseGetPaginated struct {
	exampleGetPaginatedPostgres exampleDomain.ExamplePostgresqlGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseGetPaginated(exampleGetPaginatedPostgres exampleDomain.ExamplePostgresqlGetPaginated) exampleDomain.ExampleUsecaseGetPaginated {
	return &ExampleUsecaseGetPaginated{exampleGetPaginatedPostgres: exampleGetPaginatedPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseGetPaginated) Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*exampleDomain.Example], error) {
	pagination := sharedDomain.NewPagination(page, pageSize)

	examples, total, err := uc.exampleGetPaginatedPostgres.Execute(ctx, pagination.Limit(), pagination.Offset(), search)
	if err != nil {
		return nil, err
	}

	return &sharedDomain.Page[*exampleDomain.Example]{
		Items:      examples,
		Pagination: pagination,
		Total:      total,
	}, nil
}
