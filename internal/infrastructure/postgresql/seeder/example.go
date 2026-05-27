package seeder

import (
	time "time"

	entity "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/entity"
)

func SeedExample() any {
	now := time.Now().UTC()

	return []entity.Example{
		{
			ID:          "11111111-1111-1111-1111-111111111111",
			Name:        "Example One",
			Description: "Default seeded example",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          "22222222-2222-2222-2222-222222222222",
			Name:        "Example Two",
			Description: "Another default seeded example",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
}
