package postgresql

import (
	log "log"

	seeder "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/seeder"
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
)

type seedData func() any

var seeds = []seedData{
	seeder.SeedExample,
}

func RunSeeders(db *gorm.DB) {
	for _, seed := range seeds {
		if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(seed()).Error; err != nil {
			log.Fatalf("Seeders failed to run: %v", err)
		}
	}

	log.Println("Seeders executed successfully")
}
