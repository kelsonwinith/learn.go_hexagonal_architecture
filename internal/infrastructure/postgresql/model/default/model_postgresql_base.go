package model

import (
	time "time"
)

// ============================================================================
// Types
// ============================================================================

type BaseModel struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedBy string    `gorm:"column:created_by;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedBy string    `gorm:"column:updated_by;type:varchar(255);not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}
