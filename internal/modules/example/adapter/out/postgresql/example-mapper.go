package postgresql

import (
	model "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func toExampleModel(example *exampleDomain.Example) *model.Example {
	return &model.Example{
		BaseModel: defaultModel.BaseModel{
			ID:        example.ID,
			CreatedAt: example.CreatedAt,
			UpdatedAt: example.UpdatedAt,
		},
		Name:        example.Name,
		Description: example.Description,
	}
}

func toExampleModels(examples []*exampleDomain.Example) []*model.Example {
	entities := make([]*model.Example, len(examples))
	for i, example := range examples {
		entities[i] = toExampleModel(example)
	}

	return entities
}

func toExampleDomain(entity *model.Example) *exampleDomain.Example {
	return &exampleDomain.Example{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func toExampleDomains(entities []*model.Example) []*exampleDomain.Example {
	examples := make([]*exampleDomain.Example, len(entities))
	for i, entity := range entities {
		examples[i] = toExampleDomain(entity)
	}
	return examples
}
