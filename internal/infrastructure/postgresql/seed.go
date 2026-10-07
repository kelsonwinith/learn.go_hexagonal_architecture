package postgresql

import (
	log "log"

	postgresqlSeed "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/seed"
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
)

// ============================================================================
// Functions
// ============================================================================

func allSeeds() []func() any {
	return []func() any{
		postgresqlSeed.SeedExampleBasic,
		postgresqlSeed.SeedExampleAdvanced,
	}
}

func RunSeeders(db *gorm.DB) {
	for _, seed := range allSeeds() {
		if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(seed()).Error; err != nil {
			log.Fatalf("Seeders failed to run: %v", err)
		}
	}

	log.Println("Seeders executed successfully")
}
