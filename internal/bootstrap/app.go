package bootstrap

import (
	"log"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

func Run() {
	// Config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// PostgreSQL
	db, err := postgresql.NewDBConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	defer sqlDB.Close()

	// PostgreSQL Migrations and Seeders
	postgresql.RunAutoMigrations(db)
	postgresql.RunSeeders(db)

	// Initialize Fiber App
	app := fiber.New(fiber.Config{
		StructValidator: sharedFiber.NewValidator(),
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))
	app.Get("/swagger/*", swaggo.HandlerDefault)

	// Initialize Modules
	example.Init(app, db)

	// Start Server
	log.Printf("Server listening on port %s", config.App.Port)
	if err := app.Listen(":" + config.App.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
