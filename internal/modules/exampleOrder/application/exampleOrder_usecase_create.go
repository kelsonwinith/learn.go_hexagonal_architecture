package application

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderUsecaseCreate struct {
	exampleUserModuleGetByID           exampleOrderDomain.ExampleUserModuleGetByID
	exampleProductModuleGetByID        exampleOrderDomain.ExampleProductModuleGetByID
	exampleOrderPostgresqlTransaction  exampleOrderDomain.ExampleOrderPostgresqlTransaction
	exampleOrderCreatePostgres         exampleOrderDomain.ExampleOrderPostgresqlCreate
	exampleOrderCreateProductsPostgres exampleOrderDomain.ExampleOrderProductPostgresqlCreateMultiple
	exampleOrderEventPublisher         exampleOrderDomain.ExampleOrderEventPublisher
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderUsecaseCreate(
	exampleUserModuleGetByID exampleOrderDomain.ExampleUserModuleGetByID,
	exampleProductModuleGetByID exampleOrderDomain.ExampleProductModuleGetByID,
	exampleOrderPostgresqlTransaction exampleOrderDomain.ExampleOrderPostgresqlTransaction,
	exampleOrderCreatePostgres exampleOrderDomain.ExampleOrderPostgresqlCreate,
	exampleOrderCreateProductsPostgres exampleOrderDomain.ExampleOrderProductPostgresqlCreateMultiple,
	exampleOrderEventPublisher exampleOrderDomain.ExampleOrderEventPublisher,
) exampleOrderDomain.ExampleOrderUsecaseCreate {
	return &ExampleOrderUsecaseCreate{
		exampleUserModuleGetByID:           exampleUserModuleGetByID,
		exampleProductModuleGetByID:        exampleProductModuleGetByID,
		exampleOrderPostgresqlTransaction:  exampleOrderPostgresqlTransaction,
		exampleOrderCreatePostgres:         exampleOrderCreatePostgres,
		exampleOrderCreateProductsPostgres: exampleOrderCreateProductsPostgres,
		exampleOrderEventPublisher:         exampleOrderEventPublisher,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleOrderUsecaseCreate) Execute(ctx context.Context, exampleOrderInput exampleOrderDomain.ExampleOrder) (*exampleOrderDomain.ExampleOrder, error) {
	// Cross-module: ensure the referenced user exists.
	if _, err := uc.exampleUserModuleGetByID.Execute(ctx, exampleOrderInput.UserID); err != nil {
		return nil, err
	}

	// Cross-module: resolve each referenced product and use its name.
	for _, exampleOrderProduct := range exampleOrderInput.Products {
		exampleOrderProductInfo, err := uc.exampleProductModuleGetByID.Execute(ctx, exampleOrderProduct.ProductID)
		if err != nil {
			return nil, err
		}
		exampleOrderProduct.Name = exampleOrderProductInfo.Name
	}

	exampleOrder, err := exampleOrderDomain.NewExampleOrder(exampleOrderInput.Name, exampleOrderInput.Description, exampleOrderInput.UserID, exampleOrderInput.Products, exampleOrderInput.CreatedBy)
	if err != nil {
		return nil, err
	}

	err = uc.exampleOrderPostgresqlTransaction.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := uc.exampleOrderCreatePostgres.Execute(ctx, exampleOrder); err != nil {
			return err
		}

		for _, exampleOrderProduct := range exampleOrder.Products {
			exampleOrderProduct.OrderID = exampleOrder.ID
		}

		return uc.exampleOrderCreateProductsPostgres.Execute(ctx, exampleOrder.Products)
	})
	if err != nil {
		return nil, err
	}

	exampleOrderEvent := exampleOrderDomain.NewEvent(exampleOrderDomain.EventTypeOrderCreated, exampleOrder.ID)
	if err := uc.exampleOrderEventPublisher.Execute(ctx, exampleOrderEvent); err != nil {
		return nil, err
	}

	return exampleOrder, nil
}
