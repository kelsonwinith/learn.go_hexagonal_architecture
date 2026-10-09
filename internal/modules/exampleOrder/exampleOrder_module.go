package exampleOrder

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleOrderFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/in/fiber"
	exampleOrderEventLog "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/eventlog"
	exampleOrderPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/postgresql"
	exampleOrderUseCase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/application"
	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Functions
// ============================================================================

func Init(app *fiber.App, db *gorm.DB, userReader exampleOrderDomain.ExampleOrderUserReader, productReader exampleOrderDomain.ExampleOrderProductReader) {
	// Adapters Out - PostgreSQL
	postgresql := sharedPostgresql.NewPostgresql(db)
	postgresqlTransaction := sharedPostgresql.NewPostgresqlTransaction(postgresql)

	exampleOrderCreatePostgres := exampleOrderPostgresql.NewExampleOrderPostgresqlCreate(postgresql)
	exampleOrderCreateProductsPostgres := exampleOrderPostgresql.NewExampleOrderProductPostgresqlCreateMultiple(postgresql)
	exampleOrderGetByIDPostgres := exampleOrderPostgresql.NewExampleOrderPostgresqlGetByID(postgresql)

	// Adapters Out - Event Log
	exampleOrderEventPublisher := exampleOrderEventLog.NewExampleOrderEventLogPublisher()

	// Use Cases
	exampleOrderUsecaseCreate := exampleOrderUseCase.NewExampleOrderUsecaseCreate(userReader, productReader, postgresqlTransaction, exampleOrderCreatePostgres, exampleOrderCreateProductsPostgres, exampleOrderEventPublisher)
	exampleOrderUsecaseGetByID := exampleOrderUseCase.NewExampleOrderUsecaseGetByID(exampleOrderGetByIDPostgres)

	// Adapters In - Fiber
	authMiddleware := sharedFiber.NewAuth()

	exampleOrderFiberCreate := exampleOrderFiber.NewExampleOrderFiberCreate(exampleOrderUsecaseCreate)
	exampleOrderFiberGetByID := exampleOrderFiber.NewExampleOrderFiberGetByID(exampleOrderUsecaseGetByID)

	exampleOrderRoutes := app.Group("/api/v1/exampleorder")
	exampleOrderRoutes.Post("/", authMiddleware, exampleOrderFiberCreate.Handle)
	exampleOrderRoutes.Get("/:id", authMiddleware, exampleOrderFiberGetByID.Handle)
}
