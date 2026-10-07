package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Constants
// ============================================================================

const exampleAdvancedChildTableName = "example_advanced_child"

// ============================================================================
// Types
// ============================================================================

type ExampleAdvancedChildModel struct {
	defaultModel.BaseModel
	defaultModel.DeleteModel

	ParentID string `gorm:"column:parent_id;type:uuid;not null;index"`
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	Quantity int    `gorm:"column:quantity;not null"`
}

// ============================================================================
// Methods
// ============================================================================

func (ExampleAdvancedChildModel) TableName() string {
	return exampleAdvancedChildTableName
}

// ============================================================================
// Functions
// ============================================================================

func ExampleAdvancedChildTable() string {
	return exampleAdvancedChildTableName
}
