package postgresql

import (
	fmt "fmt"
	log "log"
	time "time"

	config "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
	postgresDriver "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	schema "gorm.io/gorm/schema"
)

func NewDBConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgreSQL.DBHost, cfg.PostgreSQL.DBPort, cfg.PostgreSQL.DBUser, cfg.PostgreSQL.DBPassword, cfg.PostgreSQL.DBName, cfg.PostgreSQL.DBSSLMode)

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}
