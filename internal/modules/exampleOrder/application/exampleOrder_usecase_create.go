package application

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderUsecaseCreate struct {
	userReader             exampleOrderDomain.ExampleOrderUserReader
	productReader          exampleOrderDomain.ExampleOrderProductReader
	postgresqlTransaction  exampleOrderDomain.ExampleOrderPostgresqlTransaction
	createPostgres         exampleOrderDomain.ExampleOrderPostgresqlCreate
	createProductsPostgres exampleOrderDomain.ExampleOrderProductPostgresqlCreateMultiple
	eventPublisher         exampleOrderDomain.ExampleOrderEventPublisher
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderUsecaseCreate(
	userReader exampleOrderDomain.ExampleOrderUserReader,
	productReader exampleOrderDomain.ExampleOrderProductReader,
	postgresqlTransaction exampleOrderDomain.ExampleOrderPostgresqlTransaction,
	createPostgres exampleOrderDomain.ExampleOrderPostgresqlCreate,
	createProductsPostgres exampleOrderDomain.ExampleOrderProductPostgresqlCreateMultiple,
	eventPublisher exampleOrderDomain.ExampleOrderEventPublisher,
) exampleOrderDomain.ExampleOrderUsecaseCreate {
	return &ExampleOrderUsecaseCreate{
		userReader:             userReader,
		productReader:          productReader,
		postgresqlTransaction:  postgresqlTransaction,
		createPostgres:         createPostgres,
		createProductsPostgres: createProductsPostgres,
		eventPublisher:         eventPublisher,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleOrderUsecaseCreate) Execute(ctx context.Context, input exampleOrderDomain.ExampleOrder) (*exampleOrderDomain.ExampleOrder, error) {
	// Cross-module: ensure the referenced user exists.
	if _, err := uc.userReader.Execute(ctx, input.UserID); err != nil {
		return nil, err
	}

	// Cross-module: resolve each referenced product and use its name.
	for _, orderProduct := range input.Products {
		product, err := uc.productReader.Execute(ctx, orderProduct.ProductID)
		if err != nil {
			return nil, err
		}
		orderProduct.Name = product.Name
	}

	order, err := exampleOrderDomain.NewExampleOrder(input.Name, input.Description, input.UserID, input.Products, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	err = uc.postgresqlTransaction.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := uc.createPostgres.Execute(ctx, order); err != nil {
			return err
		}

		for _, orderProduct := range order.Products {
			orderProduct.OrderID = order.ID
		}

		return uc.createProductsPostgres.Execute(ctx, order.Products)
	})
	if err != nil {
		return nil, err
	}

	event := exampleOrderDomain.NewEvent(exampleOrderDomain.EventTypeOrderCreated, order.ID)
	if err := uc.eventPublisher.Execute(ctx, event); err != nil {
		return nil, err
	}

	return order, nil
}
