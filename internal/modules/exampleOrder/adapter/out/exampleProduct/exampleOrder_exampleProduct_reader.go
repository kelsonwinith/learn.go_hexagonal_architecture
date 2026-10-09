package exampleProduct

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderProductReader struct {
	productGetByID exampleProductDomain.ExampleProductUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderProductReader(productGetByID exampleProductDomain.ExampleProductUsecaseGetByID) *ExampleOrderProductReader {
	return &ExampleOrderProductReader{productGetByID: productGetByID}
}

// ============================================================================
// Methods
// ============================================================================

func (r *ExampleOrderProductReader) Execute(ctx context.Context, id string) (exampleOrderDomain.ExampleOrderProductInfo, error) {
	product, err := r.productGetByID.Execute(ctx, id)
	if err != nil {
		return exampleOrderDomain.ExampleOrderProductInfo{}, err
	}

	return exampleOrderDomain.ExampleOrderProductInfo{ID: product.ID, Name: product.Name, Price: product.Price}, nil
}
