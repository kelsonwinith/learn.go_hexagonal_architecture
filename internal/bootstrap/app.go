package bootstrap

import (
	errors "errors"
	log "log"
	http "net/http"
	os "os"
	signal "os/signal"
	syscall "syscall"
	time "time"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
	fiber "github.com/gofiber/fiber/v3"
	cors "github.com/gofiber/fiber/v3/middleware/cors"
	logger "github.com/gofiber/fiber/v3/middleware/logger"

	config "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql"
	exampleOrder "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder"
	exampleOrderExampleProduct "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/exampleproduct"
	exampleOrderExampleUser "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/exampleuser"
	exampleProduct "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct"
	exampleUser "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Functions
// ============================================================================

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
	postgresql.RunMigrations(db)
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
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "example-user-id"},
	}))
	app.Get("/swagger/*", swaggo.HandlerDefault)

	// Initialize Modules
	exampleUserService := exampleUser.Init(app, db)
	exampleProductService := exampleProduct.Init(app, db)

	// Cross-module adapters: exampleOrder consumes exampleUser and exampleProduct services
	exampleOrderUserReader := exampleOrderExampleUser.NewExampleOrderUserReader(exampleUserService)
	exampleOrderProductReader := exampleOrderExampleProduct.NewExampleOrderProductReader(exampleProductService)

	exampleOrder.Init(app, db, exampleOrderUserReader, exampleOrderProductReader)

	// Start Server with Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on port %s", config.App.Port)
		if err := app.Listen(":" + config.App.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server gracefully...")

	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited cleanly")
}
