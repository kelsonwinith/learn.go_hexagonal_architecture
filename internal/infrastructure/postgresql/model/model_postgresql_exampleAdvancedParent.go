package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Constants
// ============================================================================

const exampleAdvancedParentTableName = "example_advanced_parent"

// ============================================================================
// Types
// ============================================================================

type ExampleAdvancedParentModel struct {
	defaultModel.BaseModel
	defaultModel.DeleteModel

	Name        string                       `gorm:"column:name;type:varchar(255);not null"`
	Description string                       `gorm:"column:description;type:varchar(255)"`
	Children    []*ExampleAdvancedChildModel `gorm:"foreignKey:ParentID;references:ID"`
}

// ============================================================================
// Methods
// ============================================================================

func (ExampleAdvancedParentModel) TableName() string {
	return exampleAdvancedParentTableName
}

// ============================================================================
// Functions
// ============================================================================

func ExampleAdvancedParentTable() string {
	return exampleAdvancedParentTableName
}
