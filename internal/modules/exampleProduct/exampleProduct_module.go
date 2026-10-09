package exampleProduct

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleProductFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/adapter/in/fiber"
	exampleProductPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/adapter/out/postgresql"
	exampleProductUseCase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/application"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Functions
// ============================================================================

func Init(app *fiber.App, db *gorm.DB) exampleProductDomain.ExampleProductUsecaseGetByID {
	// Adapters Out - PostgreSQL
	postgresql := sharedPostgresql.NewPostgresql(db)
	postgresqlTransaction := sharedPostgresql.NewPostgresqlTransaction(postgresql)

	exampleProductPostgresqlCreate := exampleProductPostgresql.NewExampleProductPostgresqlCreate(postgresql)
	exampleProductPostgresqlGetAll := exampleProductPostgresql.NewExampleProductPostgresqlGetAll(postgresql)
	exampleProductPostgresqlGetPaginated := exampleProductPostgresql.NewExampleProductPostgresqlGetPaginated(postgresql)
	exampleProductPostgresqlGetByID := exampleProductPostgresql.NewExampleProductPostgresqlGetByID(postgresql)
	exampleProductPostgresqlUpdate := exampleProductPostgresql.NewExampleProductPostgresqlUpdate(postgresql)
	exampleProductPostgresqlDelete := exampleProductPostgresql.NewExampleProductPostgresqlDelete(postgresql)
	exampleProductPostgresqlCreateMultiple := exampleProductPostgresql.NewExampleProductPostgresqlCreateMultiple(postgresql)

	// Use Cases
	exampleProductUsecaseCreate := exampleProductUseCase.NewExampleProductUsecaseCreate(exampleProductPostgresqlCreate)
	exampleProductUsecaseGetAll := exampleProductUseCase.NewExampleProductUsecaseGetAll(exampleProductPostgresqlGetAll)
	exampleProductUsecaseGetPaginated := exampleProductUseCase.NewExampleProductUsecaseGetPaginated(exampleProductPostgresqlGetPaginated)
	exampleProductUsecaseGetByID := exampleProductUseCase.NewExampleProductUsecaseGetByID(exampleProductPostgresqlGetByID)
	exampleProductUsecaseUpdate := exampleProductUseCase.NewExampleProductUsecaseUpdate(exampleProductPostgresqlUpdate, exampleProductPostgresqlGetByID)
	exampleProductUsecaseDelete := exampleProductUseCase.NewExampleProductUsecaseDelete(exampleProductPostgresqlDelete, exampleProductPostgresqlGetByID)
	exampleProductUsecaseCreateMultiple := exampleProductUseCase.NewExampleProductUsecaseCreateMultiple(postgresqlTransaction, exampleProductPostgresqlCreateMultiple)

	// Adapters In - Fiber
	authMiddleware := sharedFiber.NewAuth()

	exampleProductFiberCreate := exampleProductFiber.NewExampleProductFiberCreate(exampleProductUsecaseCreate)
	exampleProductFiberGetAll := exampleProductFiber.NewExampleProductFiberGetAll(exampleProductUsecaseGetAll)
	exampleProductFiberGetPaginated := exampleProductFiber.NewExampleProductFiberGetPaginated(exampleProductUsecaseGetPaginated)
	exampleProductFiberGetByID := exampleProductFiber.NewExampleProductFiberGetByID(exampleProductUsecaseGetByID)
	exampleProductFiberUpdate := exampleProductFiber.NewExampleProductFiberUpdate(exampleProductUsecaseUpdate)
	exampleProductFiberDelete := exampleProductFiber.NewExampleProductFiberDelete(exampleProductUsecaseDelete)
	exampleProductFiberCreateMultiple := exampleProductFiber.NewExampleProductFiberCreateMultiple(exampleProductUsecaseCreateMultiple)

	exampleProductRoutes := app.Group("/api/v1/exampleproduct")
	exampleProductRoutes.Post("/", authMiddleware, exampleProductFiberCreate.Handle)
	exampleProductRoutes.Post("/batch", authMiddleware, exampleProductFiberCreateMultiple.Handle)
	exampleProductRoutes.Get("/", exampleProductFiberGetAll.Handle)
	exampleProductRoutes.Get("/paginated", exampleProductFiberGetPaginated.Handle)
	exampleProductRoutes.Get("/:id", exampleProductFiberGetByID.Handle)
	exampleProductRoutes.Put("/:id", authMiddleware, exampleProductFiberUpdate.Handle)
	exampleProductRoutes.Delete("/:id", authMiddleware, exampleProductFiberDelete.Handle)

	return exampleProductUsecaseGetByID
}
