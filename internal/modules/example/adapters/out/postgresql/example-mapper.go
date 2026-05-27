package postgresql

import (
	postgresqlEntity "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/entity"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func toExampleEntity(example *exampleDomain.Example) *postgresqlEntity.Example {
	return &postgresqlEntity.Example{
		ID:          example.ID,
		Name:        example.Name,
		Description: example.Description,
		CreatedAt:   example.CreatedAt,
		UpdatedAt:   example.UpdatedAt,
	}
}

func toExampleEntities(examples []*exampleDomain.Example) []*postgresqlEntity.Example {
	entities := make([]*postgresqlEntity.Example, len(examples))
	for i, example := range examples {
		entities[i] = toExampleEntity(example)
	}
	return entities
}

func toExampleDomain(entity *postgresqlEntity.Example) *exampleDomain.Example {
	return &exampleDomain.Example{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func toExampleDomains(entities []*postgresqlEntity.Example) []*exampleDomain.Example {
	examples := make([]*exampleDomain.Example, len(entities))
	for i, entity := range entities {
		examples[i] = toExampleDomain(entity)
	}
	return examples
}
