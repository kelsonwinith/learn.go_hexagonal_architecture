package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUsecaseGetPaginated struct {
	exampleGetPaginatedPostgres exampleBasicDomain.ExamplePostgresqlGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUsecaseGetPaginated(exampleGetPaginatedPostgres exampleBasicDomain.ExamplePostgresqlGetPaginated) exampleBasicDomain.ExampleUsecaseGetPaginated {
	return &ExampleUsecaseGetPaginated{exampleGetPaginatedPostgres: exampleGetPaginatedPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUsecaseGetPaginated) Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*exampleBasicDomain.Example], error) {
	pagination := sharedDomain.NewPagination(page, pageSize)

	examples, total, err := uc.exampleGetPaginatedPostgres.Execute(ctx, pagination.Limit(), pagination.Offset(), search)
	if err != nil {
		return nil, err
	}

	return &sharedDomain.Page[*exampleBasicDomain.Example]{
		Items:      examples,
		Pagination: pagination,
		Total:      total,
	}, nil
}
