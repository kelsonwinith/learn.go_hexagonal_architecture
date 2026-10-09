package exampleProduct

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleProductFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/adapter/in/fiber"
	exampleProductPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/adapter/out/postgresql"
	exampleProductUsecase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/application"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Functions
// ============================================================================

func Init(app *fiber.App, db *gorm.DB, authMiddleware fiber.Handler) exampleProductDomain.ExampleProductUsecaseGetByID {
	// Adapters Out - PostgreSQL
	postgresql := sharedPostgresql.NewPostgresql(db)

	exampleProductPostgresqlCreate := exampleProductPostgresql.NewExampleProductPostgresqlCreate(postgresql)
	exampleProductPostgresqlGetByID := exampleProductPostgresql.NewExampleProductPostgresqlGetByID(postgresql)
	exampleProductPostgresqlGetPaginated := exampleProductPostgresql.NewExampleProductPostgresqlGetPaginated(postgresql)
	exampleProductPostgresqlUpdate := exampleProductPostgresql.NewExampleProductPostgresqlUpdate(postgresql)
	exampleProductPostgresqlDelete := exampleProductPostgresql.NewExampleProductPostgresqlDelete(postgresql)

	// Use Cases
	exampleProductUsecaseCreate := exampleProductUsecase.NewExampleProductUsecaseCreate(exampleProductPostgresqlCreate)
	exampleProductUsecaseGetByID := exampleProductUsecase.NewExampleProductUsecaseGetByID(exampleProductPostgresqlGetByID)
	exampleProductUsecaseGetPaginated := exampleProductUsecase.NewExampleProductUsecaseGetPaginated(exampleProductPostgresqlGetPaginated)
	exampleProductUsecaseUpdate := exampleProductUsecase.NewExampleProductUsecaseUpdate(exampleProductPostgresqlUpdate, exampleProductPostgresqlGetByID)
	exampleProductUsecaseDelete := exampleProductUsecase.NewExampleProductUsecaseDelete(exampleProductPostgresqlDelete, exampleProductPostgresqlGetByID)

	// Adapters In - Fiber
	exampleProductFiberCreate := exampleProductFiber.NewExampleProductFiberCreate(exampleProductUsecaseCreate)
	exampleProductFiberGetByID := exampleProductFiber.NewExampleProductFiberGetByID(exampleProductUsecaseGetByID)
	exampleProductFiberGetPaginated := exampleProductFiber.NewExampleProductFiberGetPaginated(exampleProductUsecaseGetPaginated)
	exampleProductFiberUpdate := exampleProductFiber.NewExampleProductFiberUpdate(exampleProductUsecaseUpdate)
	exampleProductFiberDelete := exampleProductFiber.NewExampleProductFiberDelete(exampleProductUsecaseDelete)

	exampleProductRoutes := app.Group("/api/v1/exampleproduct")
	exampleProductRoutes.Post("/", authMiddleware, exampleProductFiberCreate.Handle)
	exampleProductRoutes.Get("/paginated", exampleProductFiberGetPaginated.Handle)
	exampleProductRoutes.Get("/:id", exampleProductFiberGetByID.Handle)
	exampleProductRoutes.Put("/:id", authMiddleware, exampleProductFiberUpdate.Handle)
	exampleProductRoutes.Delete("/:id", authMiddleware, exampleProductFiberDelete.Handle)

	return exampleProductUsecaseGetByID
}
