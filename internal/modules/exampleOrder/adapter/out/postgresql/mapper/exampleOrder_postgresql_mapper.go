package mapper

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
)

// ============================================================================
// Functions
// ============================================================================

func ToExampleOrderModel(exampleOrder *exampleOrderDomain.ExampleOrder) *postgresqlModel.ExampleOrderModel {
	return &postgresqlModel.ExampleOrderModel{
		BaseModel: defaultModel.BaseModel{
			ID:        exampleOrder.ID,
			CreatedAt: exampleOrder.CreatedAt,
			UpdatedAt: exampleOrder.UpdatedAt,
			CreatedBy: exampleOrder.CreatedBy,
			UpdatedBy: exampleOrder.UpdatedBy,
		},
		Name:        exampleOrder.Name,
		Description: exampleOrder.Description,
		UserID:      exampleOrder.UserID,
	}
}

func ToExampleOrderProductModel(exampleOrderProduct *exampleOrderDomain.ExampleOrderProduct) *postgresqlModel.ExampleOrderProductModel {
	return &postgresqlModel.ExampleOrderProductModel{
		BaseModel: defaultModel.BaseModel{
			ID:        exampleOrderProduct.ID,
			CreatedAt: exampleOrderProduct.CreatedAt,
			UpdatedAt: exampleOrderProduct.UpdatedAt,
			CreatedBy: exampleOrderProduct.CreatedBy,
			UpdatedBy: exampleOrderProduct.UpdatedBy,
		},
		OrderID:   exampleOrderProduct.OrderID,
		ProductID: exampleOrderProduct.ProductID,
		Name:      exampleOrderProduct.Name,
		Quantity:  exampleOrderProduct.Quantity,
	}
}

func ToExampleOrderProductModels(exampleOrderProducts []*exampleOrderDomain.ExampleOrderProduct) []*postgresqlModel.ExampleOrderProductModel {
	entities := make([]*postgresqlModel.ExampleOrderProductModel, len(exampleOrderProducts))
	for i, exampleOrderProduct := range exampleOrderProducts {
		entities[i] = ToExampleOrderProductModel(exampleOrderProduct)
	}

	return entities
}

func ToExampleOrderDomain(entity *postgresqlModel.ExampleOrderModel) *exampleOrderDomain.ExampleOrder {
	exampleOrder := &exampleOrderDomain.ExampleOrder{
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
		exampleOrder.Products = make([]*exampleOrderDomain.ExampleOrderProduct, len(entity.Products))
		for i, exampleOrderProduct := range entity.Products {
			exampleOrder.Products[i] = ToExampleOrderProductDomain(exampleOrderProduct)
		}
	}

	return exampleOrder
}

func ToExampleOrderDomains(entities []*postgresqlModel.ExampleOrderModel) []*exampleOrderDomain.ExampleOrder {
	exampleOrders := make([]*exampleOrderDomain.ExampleOrder, len(entities))
	for i, entity := range entities {
		exampleOrders[i] = ToExampleOrderDomain(entity)
	}

	return exampleOrders
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
