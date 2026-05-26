package postgresql

import (
	log "log"
	time "time"

	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
)

func RunSeeders(db *gorm.DB) {
	now := time.Now().UTC()
	examples := []Example{
		{
			ID:          "11111111-1111-1111-1111-111111111111",
			Name:        "Example One",
			Description: "Default seeded example",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          "22222222-2222-2222-2222-222222222222",
			Name:        "Example Two",
			Description: "Another default seeded example",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&examples).Error; err != nil {
		log.Fatalf("Seeders failed to run: %v", err)
	}

	log.Println("Seeders executed successfully")
}
