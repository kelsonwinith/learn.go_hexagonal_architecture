package mapper

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Functions
// ============================================================================

func ToExampleBasicModel(example *exampleBasicDomain.ExampleBasic) *postgresqlModel.ExampleBasicModel {
	return &postgresqlModel.ExampleBasicModel{
		BaseModel: defaultModel.BaseModel{
			ID:        example.ID,
			CreatedAt: example.CreatedAt,
			UpdatedAt: example.UpdatedAt,
			CreatedBy: example.CreatedBy,
			UpdatedBy: example.UpdatedBy,
		},
		Name:        example.Name,
		Description: example.Description,
	}
}

func ToExampleBasicModels(examples []*exampleBasicDomain.ExampleBasic) []*postgresqlModel.ExampleBasicModel {
	entities := make([]*postgresqlModel.ExampleBasicModel, len(examples))
	for i, example := range examples {
		entities[i] = ToExampleBasicModel(example)
	}

	return entities
}

func ToExampleBasicDomain(entity *postgresqlModel.ExampleBasicModel) *exampleBasicDomain.ExampleBasic {
	return &exampleBasicDomain.ExampleBasic{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedBy:   entity.CreatedBy,
		UpdatedBy:   entity.UpdatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToExampleBasicDomains(entities []*postgresqlModel.ExampleBasicModel) []*exampleBasicDomain.ExampleBasic {
	examples := make([]*exampleBasicDomain.ExampleBasic, len(entities))
	for i, entity := range entities {
		examples[i] = ToExampleBasicDomain(entity)
	}
	return examples
}
