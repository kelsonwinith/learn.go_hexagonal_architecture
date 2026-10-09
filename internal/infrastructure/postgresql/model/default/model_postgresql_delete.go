package model

import (
	gorm "gorm.io/gorm"
)

// ============================================================================
// Types
// ============================================================================

type DeleteModel struct {
	DeletedBy string         `gorm:"column:deleted_by;type:varchar(255)"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
