package bootstrap

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "github.com/kelsonwinith/learn.go-hexagonal-architecture/docs"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/env"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/migrations"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example"
)

func Run() {
	// Load Config
	cfg := env.LoadConfig()

	// Connect Database
	db, err := postgresql.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Fiber App
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Initialize Modules
	example.Init(app, db)

	// Run Migrations
	migrations.RunMigrations(cfg)

	// Start Server
	log.Printf("Server listening on port %s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
