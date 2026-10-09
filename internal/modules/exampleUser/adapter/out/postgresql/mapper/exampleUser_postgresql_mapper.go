package mapper

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Functions
// ============================================================================

func ToExampleUserModel(example *exampleUserDomain.ExampleUser) *postgresqlModel.ExampleUserModel {
	return &postgresqlModel.ExampleUserModel{
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

func ToExampleUserModels(examples []*exampleUserDomain.ExampleUser) []*postgresqlModel.ExampleUserModel {
	entities := make([]*postgresqlModel.ExampleUserModel, len(examples))
	for i, example := range examples {
		entities[i] = ToExampleUserModel(example)
	}

	return entities
}

func ToExampleUserDomain(entity *postgresqlModel.ExampleUserModel) *exampleUserDomain.ExampleUser {
	return &exampleUserDomain.ExampleUser{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedBy:   entity.CreatedBy,
		UpdatedBy:   entity.UpdatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToExampleUserDomains(entities []*postgresqlModel.ExampleUserModel) []*exampleUserDomain.ExampleUser {
	examples := make([]*exampleUserDomain.ExampleUser, len(entities))
	for i, entity := range entities {
		examples[i] = ToExampleUserDomain(entity)
	}
	return examples
}
