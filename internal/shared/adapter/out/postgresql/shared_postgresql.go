package postgresql

import (
	context "context"

	gorm "gorm.io/gorm"
)

type txKeyType struct{}

type Postgresql struct {
	DB *gorm.DB
}

var txKey = txKeyType{}

func NewPostgresql(db *gorm.DB) *Postgresql {
	return &Postgresql{
		DB: db,
	}
}

func (p *Postgresql) GetExecutor(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}
	return p.DB.WithContext(ctx)
}
