package seed

import (
	time "time"

	model "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	baseModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Functions
// ============================================================================

func SeedExampleUser() any {
	now := time.Now().UTC()

	return []model.ExampleUserModel{
		{
			BaseModel: baseModel.BaseModel{
				ID:        "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:     "Example User One",
			Email:    "user.one@example.com",
			Password: "password123",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "a188ed2e-efe8-4c20-85e3-33f8669698ab",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:     "Example User Two",
			Email:    "user.two@example.com",
			Password: "password123",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "d0594e79-55b8-438e-998a-6a2a78e8bcaa",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:     "Example User Three",
			Email:    "user.three@example.com",
			Password: "password123",
		},
	}
}
