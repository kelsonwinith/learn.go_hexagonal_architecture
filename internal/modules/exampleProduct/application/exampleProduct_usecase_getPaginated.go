package application

import (
	context "context"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductUsecaseGetPaginated struct {
	exampleGetPaginatedPostgres exampleProductDomain.ExampleProductPostgresqlGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseGetPaginated(exampleGetPaginatedPostgres exampleProductDomain.ExampleProductPostgresqlGetPaginated) exampleProductDomain.ExampleProductUsecaseGetPaginated {
	return &ExampleProductUsecaseGetPaginated{exampleGetPaginatedPostgres: exampleGetPaginatedPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseGetPaginated) Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*exampleProductDomain.ExampleProduct], error) {
	pagination := sharedDomain.NewPagination(page, pageSize)

	examples, total, err := uc.exampleGetPaginatedPostgres.Execute(ctx, pagination.Limit(), pagination.Offset(), search)
	if err != nil {
		return nil, err
	}

	return &sharedDomain.Page[*exampleProductDomain.ExampleProduct]{
		Items:      examples,
		Pagination: pagination,
		Total:      total,
	}, nil
}
