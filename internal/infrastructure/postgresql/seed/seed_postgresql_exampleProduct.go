package seed

import (
	time "time"

	model "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	baseModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Functions
// ============================================================================

func SeedExampleProduct() any {
	now := time.Now().UTC()

	return []model.ExampleProductModel{
		{
			BaseModel: baseModel.BaseModel{
				ID:        "a1b2c3d4-0001-4a2b-8c3d-000000000001",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
				UpdatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
			},
			Name:        "Product One",
			Description: "Default seeded product 01",
			Price:       1000,
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "a1b2c3d4-0002-4a2b-8c3d-000000000002",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
				UpdatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
			},
			Name:        "Product Two",
			Description: "Default seeded product 02",
			Price:       2500,
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "a1b2c3d4-0003-4a2b-8c3d-000000000003",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
				UpdatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
			},
			Name:        "Product Three",
			Description: "Default seeded product 03",
			Price:       5000,
		},
	}
}
