package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Constants
// ============================================================================

const exampleUserTableName = "example_user"

// ============================================================================
// Types
// ============================================================================

type ExampleUserModel struct {
	defaultModel.BaseModel
	defaultModel.DeleteModel

	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:varchar(255)"`
}

// ============================================================================
// Methods
// ============================================================================

func (ExampleUserModel) TableName() string {
	return exampleUserTableName
}

// ============================================================================
// Functions
// ============================================================================

func ExampleUserTable() string {
	return exampleUserTableName
}
