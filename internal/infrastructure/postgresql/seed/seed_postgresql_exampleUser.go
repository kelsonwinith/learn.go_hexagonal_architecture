package seed

import (
	time "time"

	model "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	baseModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	golangBcrypt "golang.org/x/crypto/bcrypt"
)

// ============================================================================
// Functions
// ============================================================================

func SeedExampleUser() any {
	now := time.Now().UTC()

	// All seeded users share the same password: "password123".
	hashedPassword, err := golangBcrypt.GenerateFromPassword([]byte("password123"), golangBcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return []model.ExampleUserModel{
		{
			BaseModel: baseModel.BaseModel{
				ID:        "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
				UpdatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
			},
			Name:     "Example User One",
			Email:    "user.one@example.com",
			Password: string(hashedPassword),
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "a188ed2e-efe8-4c20-85e3-33f8669698ab",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
				UpdatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
			},
			Name:     "Example User Two",
			Email:    "user.two@example.com",
			Password: string(hashedPassword),
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "d0594e79-55b8-438e-998a-6a2a78e8bcaa",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
				UpdatedBy: "df115bd0-f4c1-4d2d-ae31-c834f1ec815e",
			},
			Name:     "Example User Three",
			Email:    "user.three@example.com",
			Password: string(hashedPassword),
		},
	}
}
