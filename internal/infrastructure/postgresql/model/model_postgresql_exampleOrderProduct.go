package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Constants
// ============================================================================

const exampleOrderProductTableName = "example_order_product"

// ============================================================================
// Types
// ============================================================================

type ExampleOrderProductModel struct {
	defaultModel.BaseModel
	defaultModel.DeleteModel

	OrderID   string `gorm:"column:order_id;type:uuid;not null;index"`
	ProductID string `gorm:"column:product_id;type:uuid;not null;index"`
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	Quantity  int    `gorm:"column:quantity;not null"`
}

// ============================================================================
// Methods
// ============================================================================

func (ExampleOrderProductModel) TableName() string {
	return exampleOrderProductTableName
}

// ============================================================================
// Functions
// ============================================================================

func ExampleOrderProductTable() string {
	return exampleOrderProductTableName
}
