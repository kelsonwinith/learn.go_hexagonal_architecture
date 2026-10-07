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

	exampleBasicPostgresqlCreate := exampleBasicPostgresql.NewExampleBasicPostgresqlCreate(postgresql)
	exampleBasicPostgresqlGetAll := exampleBasicPostgresql.NewExampleBasicPostgresqlGetAll(postgresql)
	exampleBasicPostgresqlGetPaginated := exampleBasicPostgresql.NewExampleBasicPostgresqlGetPaginated(postgresql)
	exampleBasicPostgresqlGetByID := exampleBasicPostgresql.NewExampleBasicPostgresqlGetByID(postgresql)
	exampleBasicPostgresqlUpdate := exampleBasicPostgresql.NewExampleBasicPostgresqlUpdate(postgresql)
	exampleBasicPostgresqlDelete := exampleBasicPostgresql.NewExampleBasicPostgresqlDelete(postgresql)
	exampleBasicPostgresqlCreateMultiple := exampleBasicPostgresql.NewExampleBasicPostgresqlCreateMultiple(postgresql)

	// Use Cases
	exampleBasicUsecaseCreate := exampleBasicUseCase.NewExampleBasicUsecaseCreate(exampleBasicPostgresqlCreate)
	exampleBasicUsecaseGetAll := exampleBasicUseCase.NewExampleBasicUsecaseGetAll(exampleBasicPostgresqlGetAll)
	exampleBasicUsecaseGetPaginated := exampleBasicUseCase.NewExampleBasicUsecaseGetPaginated(exampleBasicPostgresqlGetPaginated)
	exampleBasicUsecaseGetByID := exampleBasicUseCase.NewExampleBasicUsecaseGetByID(exampleBasicPostgresqlGetByID)
	exampleBasicUsecaseUpdate := exampleBasicUseCase.NewExampleBasicUsecaseUpdate(exampleBasicPostgresqlUpdate, exampleBasicPostgresqlGetByID)
	exampleBasicUsecaseDelete := exampleBasicUseCase.NewExampleBasicUsecaseDelete(exampleBasicPostgresqlDelete, exampleBasicPostgresqlGetByID)
	exampleBasicUsecaseCreateMultiple := exampleBasicUseCase.NewExampleBasicUsecaseCreateMultiple(postgresqlTransaction, exampleBasicPostgresqlCreateMultiple)

	// Adapters In - Fiber
	authMiddleware := sharedFiber.NewAuth()

	exampleBasicFiberCreate := exampleBasicFiber.NewExampleBasicFiberCreate(exampleBasicUsecaseCreate)
	exampleBasicFiberGetAll := exampleBasicFiber.NewExampleBasicFiberGetAll(exampleBasicUsecaseGetAll)
	exampleBasicFiberGetPaginated := exampleBasicFiber.NewExampleBasicFiberGetPaginated(exampleBasicUsecaseGetPaginated)
	exampleBasicFiberGetByID := exampleBasicFiber.NewExampleBasicFiberGetByID(exampleBasicUsecaseGetByID)
	exampleBasicFiberUpdate := exampleBasicFiber.NewExampleBasicFiberUpdate(exampleBasicUsecaseUpdate)
	exampleBasicFiberDelete := exampleBasicFiber.NewExampleBasicFiberDelete(exampleBasicUsecaseDelete)
	exampleBasicFiberCreateMultiple := exampleBasicFiber.NewExampleBasicFiberCreateMultiple(exampleBasicUsecaseCreateMultiple)

	exampleBasicRoutes := app.Group("/api/v1/examplebasic")
	exampleBasicRoutes.Post("/", authMiddleware, exampleBasicFiberCreate.Handle)
	exampleBasicRoutes.Post("/batch", authMiddleware, exampleBasicFiberCreateMultiple.Handle)
	exampleBasicRoutes.Get("/", exampleBasicFiberGetAll.Handle)
	exampleBasicRoutes.Get("/paginated", exampleBasicFiberGetPaginated.Handle)
	exampleBasicRoutes.Get("/:id", exampleBasicFiberGetByID.Handle)
	exampleBasicRoutes.Put("/:id", authMiddleware, exampleBasicFiberUpdate.Handle)
	exampleBasicRoutes.Delete("/:id", authMiddleware, exampleBasicFiberDelete.Handle)
}
