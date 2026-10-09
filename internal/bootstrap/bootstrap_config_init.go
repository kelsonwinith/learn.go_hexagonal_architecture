package bootstrap

import (
	log "log"

	config "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
)

// ============================================================================
// Functions
// ============================================================================

func InitConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	return cfg
}
