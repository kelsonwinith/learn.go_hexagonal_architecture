package postgresql

import (
	time "time"
)

type Example struct {
	ID          string    `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	Description string    `gorm:"column:description;type:text"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null"`
}
