package seed

import (
	time "time"

	model "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	baseModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Functions
// ============================================================================

func SeedExampleOrder() any {
	now := time.Now().UTC()

	orderID := "f1e2d3c4-0001-4a2b-8c3d-000000000001"

	return []model.ExampleOrderModel{
		{
			BaseModel: baseModel.BaseModel{
				ID:        orderID,
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Order One",
			Description: "Default seeded example order 01",
			UserID:      "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
			Products: []*model.ExampleOrderProductModel{
				{
					BaseModel: baseModel.BaseModel{
						ID:        "f1e2d3c4-0002-4a2b-8c3d-000000000002",
						CreatedAt: now,
						UpdatedAt: now,
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					OrderID:   orderID,
					ProductID: "a1b2c3d4-0001-4a2b-8c3d-000000000001",
					Name:      "Example Order Product One",
					Quantity:  2,
				},
				{
					BaseModel: baseModel.BaseModel{
						ID:        "f1e2d3c4-0003-4a2b-8c3d-000000000003",
						CreatedAt: now,
						UpdatedAt: now,
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					OrderID:   orderID,
					ProductID: "a1b2c3d4-0002-4a2b-8c3d-000000000002",
					Name:      "Example Order Product Two",
					Quantity:  5,
				},
			},
		},
	}
}
