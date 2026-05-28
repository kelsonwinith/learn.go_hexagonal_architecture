package seed

import (
	time "time"

	model "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	baseModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

func SeedExample() any {
	now := time.Now().UTC()

	return []model.Example{
		{
			BaseModel: baseModel.BaseModel{
				ID:        "11111111-1111-1111-1111-111111111111",
				CreatedAt: now,
				UpdatedAt: now,
			},
			Name:        "Example One",
			Description: "Default seeded example",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "22222222-2222-2222-2222-222222222222",
				CreatedAt: now,
				UpdatedAt: now,
			},
			Name:        "Example Two",
			Description: "Another default seeded example",
		},
	}
}
