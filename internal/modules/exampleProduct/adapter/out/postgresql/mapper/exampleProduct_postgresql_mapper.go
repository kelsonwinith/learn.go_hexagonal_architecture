package mapper

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Functions
// ============================================================================

func ToExampleProductModel(exampleProduct *exampleProductDomain.ExampleProduct) *postgresqlModel.ExampleProductModel {
	return &postgresqlModel.ExampleProductModel{
		BaseModel: defaultModel.BaseModel{
			ID:        exampleProduct.ID,
			CreatedAt: exampleProduct.CreatedAt,
			UpdatedAt: exampleProduct.UpdatedAt,
			CreatedBy: exampleProduct.CreatedBy,
			UpdatedBy: exampleProduct.UpdatedBy,
		},
		Name:        exampleProduct.Name,
		Description: exampleProduct.Description,
		Price:       exampleProduct.Price,
	}
}

func ToExampleProductModels(exampleProducts []*exampleProductDomain.ExampleProduct) []*postgresqlModel.ExampleProductModel {
	entities := make([]*postgresqlModel.ExampleProductModel, len(exampleProducts))
	for i, exampleProduct := range exampleProducts {
		entities[i] = ToExampleProductModel(exampleProduct)
	}

	return entities
}

func ToExampleProductDomain(entity *postgresqlModel.ExampleProductModel) *exampleProductDomain.ExampleProduct {
	return &exampleProductDomain.ExampleProduct{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
		CreatedBy:   entity.CreatedBy,
		UpdatedBy:   entity.UpdatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToExampleProductDomains(entities []*postgresqlModel.ExampleProductModel) []*exampleProductDomain.ExampleProduct {
	exampleProducts := make([]*exampleProductDomain.ExampleProduct, len(entities))
	for i, entity := range entities {
		exampleProducts[i] = ToExampleProductDomain(entity)
	}
	return exampleProducts
}
