package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Constants
// ============================================================================

const exampleProductTableName = "example_product"

// ============================================================================
// Types
// ============================================================================

type ExampleProductModel struct {
	defaultModel.BaseModel
	defaultModel.DeleteModel

	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:varchar(255)"`
	Price       int64  `gorm:"column:price;not null;default:0"`
}

// ============================================================================
// Methods
// ============================================================================

func (ExampleProductModel) TableName() string {
	return exampleProductTableName
}

// ============================================================================
// Functions
// ============================================================================

func ExampleProductTable() string {
	return exampleProductTableName
}
