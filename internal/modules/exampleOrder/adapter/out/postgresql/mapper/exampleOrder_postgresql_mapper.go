package mapper

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
)

// ============================================================================
// Functions
// ============================================================================

func ToExampleOrderModel(order *exampleOrderDomain.ExampleOrder) *postgresqlModel.ExampleOrderModel {
	return &postgresqlModel.ExampleOrderModel{
		BaseModel: defaultModel.BaseModel{
			ID:        order.ID,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			CreatedBy: order.CreatedBy,
			UpdatedBy: order.UpdatedBy,
		},
		Name:        order.Name,
		Description: order.Description,
		UserID:      order.UserID,
	}
}

func ToExampleOrderProductModel(orderProduct *exampleOrderDomain.ExampleOrderProduct) *postgresqlModel.ExampleOrderProductModel {
	return &postgresqlModel.ExampleOrderProductModel{
		BaseModel: defaultModel.BaseModel{
			ID:        orderProduct.ID,
			CreatedAt: orderProduct.CreatedAt,
			UpdatedAt: orderProduct.UpdatedAt,
			CreatedBy: orderProduct.CreatedBy,
			UpdatedBy: orderProduct.UpdatedBy,
		},
		OrderID:   orderProduct.OrderID,
		ProductID: orderProduct.ProductID,
		Name:      orderProduct.Name,
		Quantity:  orderProduct.Quantity,
	}
}

func ToExampleOrderProductModels(products []*exampleOrderDomain.ExampleOrderProduct) []*postgresqlModel.ExampleOrderProductModel {
	entities := make([]*postgresqlModel.ExampleOrderProductModel, len(products))
	for i, orderProduct := range products {
		entities[i] = ToExampleOrderProductModel(orderProduct)
	}

	return entities
}

func ToExampleOrderDomain(entity *postgresqlModel.ExampleOrderModel) *exampleOrderDomain.ExampleOrder {
	order := &exampleOrderDomain.ExampleOrder{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		UserID:      entity.UserID,
		CreatedBy:   entity.CreatedBy,
		UpdatedBy:   entity.UpdatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}

	if entity.Products != nil {
		order.Products = make([]*exampleOrderDomain.ExampleOrderProduct, len(entity.Products))
		for i, orderProduct := range entity.Products {
			order.Products[i] = ToExampleOrderProductDomain(orderProduct)
		}
	}

	return order
}

func ToExampleOrderProductDomain(entity *postgresqlModel.ExampleOrderProductModel) *exampleOrderDomain.ExampleOrderProduct {
	return &exampleOrderDomain.ExampleOrderProduct{
		ID:        entity.ID,
		OrderID:   entity.OrderID,
		ProductID: entity.ProductID,
		Name:      entity.Name,
		Quantity:  entity.Quantity,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
