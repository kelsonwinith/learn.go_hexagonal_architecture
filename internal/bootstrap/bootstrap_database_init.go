package bootstrap

import (
	log "log"

	config "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Functions
// ============================================================================

func InitDatabase(cfg *config.Config) *gorm.DB {
	db, err := postgresql.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	postgresql.RunMigrations(db)
	postgresql.RunSeeders(db)

	return db
}

func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
}
