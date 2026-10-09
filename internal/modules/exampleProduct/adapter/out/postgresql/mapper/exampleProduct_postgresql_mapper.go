package mapper

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Functions
// ============================================================================

func ToExampleProductModel(example *exampleProductDomain.ExampleProduct) *postgresqlModel.ExampleProductModel {
	return &postgresqlModel.ExampleProductModel{
		BaseModel: defaultModel.BaseModel{
			ID:        example.ID,
			CreatedAt: example.CreatedAt,
			UpdatedAt: example.UpdatedAt,
			CreatedBy: example.CreatedBy,
			UpdatedBy: example.UpdatedBy,
		},
		Name:        example.Name,
		Description: example.Description,
		Price:       example.Price,
	}
}

func ToExampleProductModels(examples []*exampleProductDomain.ExampleProduct) []*postgresqlModel.ExampleProductModel {
	entities := make([]*postgresqlModel.ExampleProductModel, len(examples))
	for i, example := range examples {
		entities[i] = ToExampleProductModel(example)
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
	examples := make([]*exampleProductDomain.ExampleProduct, len(entities))
	for i, entity := range entities {
		examples[i] = ToExampleProductDomain(entity)
	}
	return examples
}
