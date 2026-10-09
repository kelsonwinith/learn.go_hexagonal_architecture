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
	exampleProductGetPaginatedPostgres exampleProductDomain.ExampleProductPostgresqlGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductUsecaseGetPaginated(exampleProductGetPaginatedPostgres exampleProductDomain.ExampleProductPostgresqlGetPaginated) exampleProductDomain.ExampleProductUsecaseGetPaginated {
	return &ExampleProductUsecaseGetPaginated{exampleProductGetPaginatedPostgres: exampleProductGetPaginatedPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleProductUsecaseGetPaginated) Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*exampleProductDomain.ExampleProduct], error) {
	pagination := sharedDomain.NewPagination(page, pageSize)

	exampleProducts, total, err := uc.exampleProductGetPaginatedPostgres.Execute(ctx, pagination.Limit(), pagination.Offset(), search)
	if err != nil {
		return nil, err
	}

	return &sharedDomain.Page[*exampleProductDomain.ExampleProduct]{
		Items:      exampleProducts,
		Pagination: pagination,
		Total:      total,
	}, nil
}
