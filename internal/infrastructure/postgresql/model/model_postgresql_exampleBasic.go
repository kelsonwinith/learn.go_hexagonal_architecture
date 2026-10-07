package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Constants
// ============================================================================

const exampleBasicTableName = "example_basic"

// ============================================================================
// Types
// ============================================================================

type ExampleBasicModel struct {
	defaultModel.BaseModel
	defaultModel.DeleteModel

	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:varchar(255)"`
}

// ============================================================================
// Methods
// ============================================================================

func (ExampleBasicModel) TableName() string {
	return exampleBasicTableName
}

// ============================================================================
// Functions
// ============================================================================

func ExampleBasicTable() string {
	return exampleBasicTableName
}
