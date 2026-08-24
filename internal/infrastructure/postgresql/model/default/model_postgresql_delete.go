package model

import (
	gorm "gorm.io/gorm"
)

type DeleteModel struct {
	DeletedBy int64          `gorm:"column:deleted_by;type:bigint"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
