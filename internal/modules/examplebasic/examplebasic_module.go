package exampleBasic

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleBasicFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/adapter/in/fiber"
	exampleBasicPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/adapter/out/postgresql"
	exampleBasicUseCase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/application"
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

	examplePostgresqlCreate := exampleBasicPostgresql.NewExamplePostgresqlCreate(postgresql)
	examplePostgresqlGetAll := exampleBasicPostgresql.NewExamplePostgresqlGetAll(postgresql)
	examplePostgresqlGetPaginated := exampleBasicPostgresql.NewExamplePostgresqlGetPaginated(postgresql)
	examplePostgresqlGetByID := exampleBasicPostgresql.NewExamplePostgresqlGetByID(postgresql)
	examplePostgresqlUpdate := exampleBasicPostgresql.NewExamplePostgresqlUpdate(postgresql)
	examplePostgresqlDelete := exampleBasicPostgresql.NewExamplePostgresqlDelete(postgresql)
	examplePostgresqlCreateMultiple := exampleBasicPostgresql.NewExamplePostgresqlCreateMultiple(postgresql)

	// Use Cases
	exampleUsecaseCreate := exampleBasicUseCase.NewExampleUsecaseCreate(examplePostgresqlCreate)
	exampleUsecaseGetAll := exampleBasicUseCase.NewExampleUsecaseGetAll(examplePostgresqlGetAll)
	exampleUsecaseGetPaginated := exampleBasicUseCase.NewExampleUsecaseGetPaginated(examplePostgresqlGetPaginated)
	exampleUsecaseGetByID := exampleBasicUseCase.NewExampleUsecaseGetByID(examplePostgresqlGetByID)
	exampleUsecaseUpdate := exampleBasicUseCase.NewExampleUsecaseUpdate(examplePostgresqlUpdate, examplePostgresqlGetByID)
	exampleUsecaseDelete := exampleBasicUseCase.NewExampleUsecaseDelete(examplePostgresqlDelete, examplePostgresqlGetByID)
	exampleUsecaseCreateMultiple := exampleBasicUseCase.NewExampleUsecaseCreateMultiple(postgresqlTransaction, examplePostgresqlCreateMultiple)

	// Adapters In - Fiber
	authMiddleware := sharedFiber.NewAuth()

	exampleFiberCreate := exampleBasicFiber.NewExampleFiberCreate(exampleUsecaseCreate)
	exampleFiberGetAll := exampleBasicFiber.NewExampleFiberGetAll(exampleUsecaseGetAll)
	exampleFiberGetPaginated := exampleBasicFiber.NewExampleFiberGetPaginated(exampleUsecaseGetPaginated)
	exampleFiberGetByID := exampleBasicFiber.NewExampleFiberGetByID(exampleUsecaseGetByID)
	exampleFiberUpdate := exampleBasicFiber.NewExampleFiberUpdate(exampleUsecaseUpdate)
	exampleFiberDelete := exampleBasicFiber.NewExampleFiberDelete(exampleUsecaseDelete)
	exampleFiberCreateMultiple := exampleBasicFiber.NewExampleFiberCreateMultiple(exampleUsecaseCreateMultiple)

	routes := app.Group("/api/v1/examplebasic")
	routes.Post("/", authMiddleware, exampleFiberCreate.Handle)
	routes.Post("/batch", authMiddleware, exampleFiberCreateMultiple.Handle)
	routes.Get("/", exampleFiberGetAll.Handle)
	routes.Get("/paginated", exampleFiberGetPaginated.Handle)
	routes.Get("/:id", exampleFiberGetByID.Handle)
	routes.Put("/:id", authMiddleware, exampleFiberUpdate.Handle)
	routes.Delete("/:id", authMiddleware, exampleFiberDelete.Handle)
}
