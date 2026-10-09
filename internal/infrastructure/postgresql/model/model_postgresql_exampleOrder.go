package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

// ============================================================================
// Constants
// ============================================================================

const exampleOrderTableName = "example_order"

// ============================================================================
// Types
// ============================================================================

type ExampleOrderModel struct {
	defaultModel.BaseModel
	defaultModel.DeleteModel

	Name        string                      `gorm:"column:name;type:varchar(255);not null"`
	Description string                      `gorm:"column:description;type:varchar(255)"`
	UserID      string                      `gorm:"column:user_id;type:uuid;not null;index"`
	Products    []*ExampleOrderProductModel `gorm:"foreignKey:OrderID;references:ID"`
}

// ============================================================================
// Methods
// ============================================================================

func (ExampleOrderModel) TableName() string {
	return exampleOrderTableName
}

// ============================================================================
// Functions
// ============================================================================

func ExampleOrderTable() string {
	return exampleOrderTableName
}
