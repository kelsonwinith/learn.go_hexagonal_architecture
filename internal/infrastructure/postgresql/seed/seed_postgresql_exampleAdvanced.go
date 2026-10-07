package seed

import (
	time "time"

	model "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	baseModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Functions
// ============================================================================

func SeedExampleAdvanced() any {
	now := time.Now().UTC()

	parentID := "f1e2d3c4-0001-4a2b-8c3d-000000000001"

	return []model.ExampleAdvancedParentModel{
		{
			BaseModel: baseModel.BaseModel{
				ID:        parentID,
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Advanced Parent One",
			Description: "Default seeded example advanced parent 01",
			Children: []*model.ExampleAdvancedChildModel{
				{
					BaseModel: baseModel.BaseModel{
						ID:        "f1e2d3c4-0002-4a2b-8c3d-000000000002",
						CreatedAt: now,
						UpdatedAt: now,
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					ParentID: parentID,
					Name:     "Example Advanced Child One",
					Quantity: 2,
				},
				{
					BaseModel: baseModel.BaseModel{
						ID:        "f1e2d3c4-0003-4a2b-8c3d-000000000003",
						CreatedAt: now,
						UpdatedAt: now,
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					ParentID: parentID,
					Name:     "Example Advanced Child Two",
					Quantity: 5,
				},
			},
		},
	}
}
