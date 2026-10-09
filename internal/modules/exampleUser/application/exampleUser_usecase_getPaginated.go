package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseGetPaginated struct {
	exampleGetPaginatedPostgres exampleUserDomain.ExampleUserPostgresqlGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseGetPaginated(exampleGetPaginatedPostgres exampleUserDomain.ExampleUserPostgresqlGetPaginated) exampleUserDomain.ExampleUserUsecaseGetPaginated {
	return &ExampleUserUsecaseGetPaginated{exampleGetPaginatedPostgres: exampleGetPaginatedPostgres}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseGetPaginated) Execute(ctx context.Context, page, pageSize int, search string) (*sharedDomain.Page[*exampleUserDomain.ExampleUser], error) {
	pagination := sharedDomain.NewPagination(page, pageSize)

	examples, total, err := uc.exampleGetPaginatedPostgres.Execute(ctx, pagination.Limit(), pagination.Offset(), search)
	if err != nil {
		return nil, err
	}

	return &sharedDomain.Page[*exampleUserDomain.ExampleUser]{
		Items:      examples,
		Pagination: pagination,
		Total:      total,
	}, nil
}
