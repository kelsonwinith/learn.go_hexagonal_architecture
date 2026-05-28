package model

import (
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
)

const exampleTableName = "example"

type Example struct {
	defaultModel.BaseModel

	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:text"`
}

func (Example) TableName() string {
	return exampleTableName
}

func ExampleTable() string {
	return exampleTableName
}
