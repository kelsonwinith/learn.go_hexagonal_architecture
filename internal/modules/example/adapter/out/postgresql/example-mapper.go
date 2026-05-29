package postgresql

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func toExampleModel(example *exampleDomain.Example) *postgresqlModel.ExampleModel {
	return &postgresqlModel.ExampleModel{
		BaseModel: defaultModel.BaseModel{
			ID:        example.ID,
			CreatedAt: example.CreatedAt,
			UpdatedAt: example.UpdatedAt,
		},
		Name:        example.Name,
		Description: example.Description,
	}
}

func toExampleModels(examples []*exampleDomain.Example) []*postgresqlModel.ExampleModel {
	entities := make([]*postgresqlModel.ExampleModel, len(examples))
	for i, example := range examples {
		entities[i] = toExampleModel(example)
	}

	return entities
}

func toExampleDomain(entity *postgresqlModel.ExampleModel) *exampleDomain.Example {
	return &exampleDomain.Example{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func toExampleDomains(entities []*postgresqlModel.ExampleModel) []*exampleDomain.Example {
	examples := make([]*exampleDomain.Example, len(entities))
	for i, entity := range entities {
		examples[i] = toExampleDomain(entity)
	}
	return examples
}
