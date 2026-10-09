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
			Name:        "Example One",
			Description: "Default seeded example 01",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "a188ed2e-efe8-4c20-85e3-33f8669698ab",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Two",
			Description: "Default seeded example 02",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "d0594e79-55b8-438e-998a-6a2a78e8bcaa",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Three",
			Description: "Default seeded example 03",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "6ce5738e-c00d-4fc0-b531-f649b08ff3b7",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Four",
			Description: "Default seeded example 04",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "e26d4e5b-8859-4081-ac21-2256bbf4dd18",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Five",
			Description: "Default seeded example 05",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "c3a60209-7a1a-43c7-ae08-bf0cdae26a97",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Six",
			Description: "Default seeded example 06",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "d662a80d-fcd3-45c6-adc0-461d8b11fd89",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Seven",
			Description: "Default seeded example 07",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "86554b62-7ea2-4526-8ff7-9795f3066dc3",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Eight",
			Description: "Default seeded example 08",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "eaeb188e-ec17-4a16-9f5d-1ef961c3a49f",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Nine",
			Description: "Default seeded example 09",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "b653a407-0f73-4b73-afd9-2fbfaaef338e",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Ten",
			Description: "Default seeded example 10",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "0ce1e143-581b-4189-8d82-49f0cfc3f3c8",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Eleven",
			Description: "Default seeded example 11",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "ad1eca51-cc73-4ede-a29a-a86573745f4a",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Twelve",
			Description: "Default seeded example 12",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "21b706d7-4733-48f6-8148-17381f61b3e1",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Thirteen",
			Description: "Default seeded example 13",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "b68086bb-253f-4ab7-afb4-72c5bda2f2a6",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Fourteen",
			Description: "Default seeded example 14",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "beb79e2e-3e82-48ef-b22a-87571d85b86a",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Fifteen",
			Description: "Default seeded example 15",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "cfa1b8cd-71d7-4704-b625-27250c8a690c",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Sixteen",
			Description: "Default seeded example 16",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "1814f3e8-8c77-495f-907c-1613a63fbbf6",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Seventeen",
			Description: "Default seeded example 17",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "ac86629f-d716-4106-9480-76caf3175499",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Eighteen",
			Description: "Default seeded example 18",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "c25dcbf2-33ed-477f-8fe9-e7908fe1f302",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Nineteen",
			Description: "Default seeded example 19",
		},
		{
			BaseModel: baseModel.BaseModel{
				ID:        "62e3805d-a0da-4e05-93ee-5250f8891f28",
				CreatedAt: now,
				UpdatedAt: now,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Name:        "Example Twenty",
			Description: "Default seeded example 20",
		},
	}
}
