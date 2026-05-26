package postgresql

import (
	log "log"

	gorm "gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Fatalf("UUID extension migration failed to run: %v", err)
	}

	if err := db.AutoMigrate(&Example{}); err != nil {
		log.Fatalf("Auto migration failed to run: %v", err)
	}

	log.Println("Auto migrations executed successfully")
}
