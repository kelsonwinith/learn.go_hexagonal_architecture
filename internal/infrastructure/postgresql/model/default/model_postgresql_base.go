package model

import (
	time "time"
)

type BaseModel struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}
