package bootstrap

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/migrations"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example"
)

func Run() {
	// Load Config
	config := config.LoadConfig()

	// Connect Database
	db, err := postgresql.NewDBConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Fiber App
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Initialize Modules
	example.Init(app, db)

	// Run Migrations
	migrations.RunMigrations(config)

	// Start Server
	log.Printf("Server listening on port %s", config.AppPort)
	if err := app.Listen(":" + config.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
