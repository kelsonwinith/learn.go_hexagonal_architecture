package postgresql

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
)

func RunMigrations(config *config.Config) {
	dbURL := "postgres://" + config.PostgreSQL.DBUser + ":" + config.PostgreSQL.DBPassword + "@" + config.PostgreSQL.DBHost + ":" + config.PostgreSQL.DBPort + "/" + config.PostgreSQL.DBName + "?sslmode=" + config.PostgreSQL.DBSSLMode

	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		log.Fatalf("Migration failed to initialize: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed to run: %v", err)
	}

	log.Println("Migrations executed successfully")
}
