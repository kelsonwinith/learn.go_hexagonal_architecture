package application

import (
	context "context"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicUsecaseGetPaginated struct {
	exampleGetPaginatedPostgres exampleBasicDomain.ExampleBasicPostgresqlGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicUsecaseGetPaginated(exampleGetPaginatedPostgres exampleBasicDomain.ExampleBasicPostgresqlGetPaginated) exampleBasicDomain.ExampleBasicUsecaseGetPaginated {
	return &ExampleBasicUsecaseGetPaginated{exampleGetPaginatedPostgres: exampleGetPaginatedPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleBasicUsecaseGetPaginated) Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*exampleBasicDomain.ExampleBasic], error) {
	pagination := sharedDomain.NewPagination(page, pageSize)

	examples, total, err := uc.exampleGetPaginatedPostgres.Execute(ctx, pagination.Limit(), pagination.Offset(), search)
	if err != nil {
		return nil, err
	}

	return &sharedDomain.Page[*exampleBasicDomain.ExampleBasic]{
		Items:      examples,
		Pagination: pagination,
		Total:      total,
	}, nil
}
