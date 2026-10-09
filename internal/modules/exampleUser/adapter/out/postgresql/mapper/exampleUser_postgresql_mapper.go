package mapper

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Functions
// ============================================================================

func ToExampleUserModel(exampleUser *exampleUserDomain.ExampleUser) *postgresqlModel.ExampleUserModel {
	return &postgresqlModel.ExampleUserModel{
		BaseModel: defaultModel.BaseModel{
			ID:        exampleUser.ID,
			CreatedAt: exampleUser.CreatedAt,
			UpdatedAt: exampleUser.UpdatedAt,
			CreatedBy: exampleUser.CreatedBy,
			UpdatedBy: exampleUser.UpdatedBy,
		},
		Name:     exampleUser.Name,
		Email:    exampleUser.Email,
		Password: exampleUser.Password,
	}
}

func ToExampleUserDomain(entity *postgresqlModel.ExampleUserModel) *exampleUserDomain.ExampleUser {
	return &exampleUserDomain.ExampleUser{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
