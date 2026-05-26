package postgresql

import (
	infrastructurePostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func toEntity(example *exampleDomain.Example) *infrastructurePostgresql.Example {
	return &infrastructurePostgresql.Example{
		ID:          example.ID,
		Name:        example.Name,
		Description: example.Description,
		CreatedAt:   example.CreatedAt,
		UpdatedAt:   example.UpdatedAt,
	}
}

func toEntities(examples []*exampleDomain.Example) []*infrastructurePostgresql.Example {
	entities := make([]*infrastructurePostgresql.Example, len(examples))
	for i, example := range examples {
		entities[i] = toEntity(example)
	}
	return entities
}

func toDomain(entity *infrastructurePostgresql.Example) *exampleDomain.Example {
	return &exampleDomain.Example{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func toDomains(entities []*infrastructurePostgresql.Example) []*exampleDomain.Example {
	examples := make([]*exampleDomain.Example, len(entities))
	for i, entity := range entities {
		examples[i] = toDomain(entity)
	}
	return examples
}
