package model

import (
	time "time"
)

// ============================================================================
// Types
// ============================================================================

type BaseModel struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedBy int64     `gorm:"column:created_by;type:bigint;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedBy int64     `gorm:"column:updated_by;type:bigint;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}
