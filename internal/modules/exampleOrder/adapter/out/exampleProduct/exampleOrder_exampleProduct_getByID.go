package exampleProduct

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductModuleGetByID struct {
	productGetByID exampleProductDomain.ExampleProductUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductModuleGetByID(productGetByID exampleProductDomain.ExampleProductUsecaseGetByID) *ExampleProductModuleGetByID {
	return &ExampleProductModuleGetByID{productGetByID: productGetByID}
}

// ============================================================================
// Methods
// ============================================================================

func (r *ExampleProductModuleGetByID) Execute(ctx context.Context, id string) (exampleOrderDomain.ExampleOrderProductInfo, error) {
	product, err := r.productGetByID.Execute(ctx, id)
	if err != nil {
		return exampleOrderDomain.ExampleOrderProductInfo{}, err
	}

	return exampleOrderDomain.ExampleOrderProductInfo{ID: product.ID, Name: product.Name, Price: product.Price}, nil
}
