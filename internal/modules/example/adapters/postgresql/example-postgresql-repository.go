package postgresql

import (
	"github.com/jmoiron/sqlx"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleRepository struct {
	db *sqlx.DB
}

func NewExampleRepository(db *sqlx.DB) domain.ExampleRepository {
	return &ExampleRepository{db: db}
}
