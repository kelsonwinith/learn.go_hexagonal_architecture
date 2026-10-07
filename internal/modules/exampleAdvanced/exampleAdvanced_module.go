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

	exampleAdvancedCreatePostgres := exampleAdvancedPostgresql.NewExampleAdvancedPostgresqlCreate(postgresql)
	exampleAdvancedCreateChildrenPostgres := exampleAdvancedPostgresql.NewExampleAdvancedChildPostgresqlCreateMultiple(postgresql)
	exampleAdvancedGetByIDPostgres := exampleAdvancedPostgresql.NewExampleAdvancedPostgresqlGetByID(postgresql)

	// Adapters Out - Event Log
	exampleAdvancedEventPublisher := exampleAdvancedEventLog.NewExampleAdvancedEventLogPublisher()

	// Use Cases
	exampleAdvancedUsecaseCreate := exampleAdvancedUseCase.NewExampleAdvancedUsecaseCreate(postgresqlTransaction, exampleAdvancedCreatePostgres, exampleAdvancedCreateChildrenPostgres, exampleAdvancedEventPublisher)
	exampleAdvancedUsecaseGetByID := exampleAdvancedUseCase.NewExampleAdvancedUsecaseGetByID(exampleAdvancedGetByIDPostgres)

	// Adapters In - Fiber
	authMiddleware := sharedFiber.NewAuth()

	exampleAdvancedFiberCreate := exampleAdvancedFiber.NewExampleAdvancedFiberCreate(exampleAdvancedUsecaseCreate)
	exampleAdvancedFiberGetByID := exampleAdvancedFiber.NewExampleAdvancedFiberGetByID(exampleAdvancedUsecaseGetByID)

	exampleAdvancedRoutes := app.Group("/api/v1/exampleadvanced")
	exampleAdvancedRoutes.Post("/", authMiddleware, exampleAdvancedFiberCreate.Handle)
	exampleAdvancedRoutes.Get("/:id", authMiddleware, exampleAdvancedFiberGetByID.Handle)
}
