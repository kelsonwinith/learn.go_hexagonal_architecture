package application

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderUsecaseGetPaginated struct {
	exampleOrderGetPaginatedPostgres exampleOrderDomain.ExampleOrderPostgresqlGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderUsecaseGetPaginated(exampleOrderGetPaginatedPostgres exampleOrderDomain.ExampleOrderPostgresqlGetPaginated) exampleOrderDomain.ExampleOrderUsecaseGetPaginated {
	return &ExampleOrderUsecaseGetPaginated{exampleOrderGetPaginatedPostgres: exampleOrderGetPaginatedPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleOrderUsecaseGetPaginated) Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*exampleOrderDomain.ExampleOrder], error) {
	pagination := sharedDomain.NewPagination(page, pageSize)

	orders, total, err := uc.exampleOrderGetPaginatedPostgres.Execute(ctx, pagination.Limit(), pagination.Offset(), search)
	if err != nil {
		return nil, err
	}

	return &sharedDomain.Page[*exampleOrderDomain.ExampleOrder]{
		Items:      orders,
		Pagination: pagination,
		Total:      total,
	}, nil
}
