package exampleAdvanced

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleAdvancedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/adapter/in/fiber"
	exampleAdvancedEventLog "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/adapter/out/eventlog"
	exampleAdvancedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/adapter/out/postgresql"
	exampleAdvancedUseCase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/application"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Functions
// ============================================================================

func Init(app *fiber.App, db *gorm.DB) {
	// Adapters Out - PostgreSQL
	postgresql := sharedPostgresql.NewPostgresql(db)
	postgresqlTransaction := sharedPostgresql.NewPostgresqlTransaction(postgresql)

	createPostgres := exampleAdvancedPostgresql.NewExampleAdvancedPostgresqlCreate(postgresql)
	createChildrenPostgres := exampleAdvancedPostgresql.NewExampleAdvancedChildPostgresqlCreateMultiple(postgresql)
	getByIDPostgres := exampleAdvancedPostgresql.NewExampleAdvancedPostgresqlGetByID(postgresql)

	// Adapters Out - Event Log
	eventPublisher := exampleAdvancedEventLog.NewExampleAdvancedEventLogPublisher()

	// Use Cases
	advancedUsecaseCreate := exampleAdvancedUseCase.NewExampleAdvancedUsecaseCreate(postgresqlTransaction, createPostgres, createChildrenPostgres, eventPublisher)
	advancedUsecaseGetByID := exampleAdvancedUseCase.NewExampleAdvancedUsecaseGetByID(getByIDPostgres)

	// Adapters In - Fiber
	authMiddleware := sharedFiber.NewAuth()

	advancedFiberCreate := exampleAdvancedFiber.NewExampleAdvancedFiberCreate(advancedUsecaseCreate)
	advancedFiberGetByID := exampleAdvancedFiber.NewExampleAdvancedFiberGetByID(advancedUsecaseGetByID)

	routes := app.Group("/api/v1/exampleadvanced")
	routes.Post("/", authMiddleware, advancedFiberCreate.Handle)
	routes.Get("/:id", advancedFiberGetByID.Handle)
}
