package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Constants
// ============================================================================

const exampleTableName = "example"

// ============================================================================
// Types
// ============================================================================

type ExampleModel struct {
	defaultModel.BaseModel
	defaultModel.DeleteModel

	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:text"`
}

// ============================================================================
// Methods
// ============================================================================

func (ExampleModel) TableName() string {
	return exampleTableName
}

// ============================================================================
// Functions
// ============================================================================

func ExampleTable() string {
	return exampleTableName
}
