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

	exampleProductPostgresqlCreate := exampleProductPostgresql.NewExampleProductPostgresqlCreate(postgresql)
	exampleProductPostgresqlGetByID := exampleProductPostgresql.NewExampleProductPostgresqlGetByID(postgresql)
	exampleProductPostgresqlGetPaginated := exampleProductPostgresql.NewExampleProductPostgresqlGetPaginated(postgresql)
	exampleProductPostgresqlUpdate := exampleProductPostgresql.NewExampleProductPostgresqlUpdate(postgresql)
	exampleProductPostgresqlDelete := exampleProductPostgresql.NewExampleProductPostgresqlDelete(postgresql)

	// Use Cases
	exampleProductUsecaseCreate := exampleProductUseCase.NewExampleProductUsecaseCreate(exampleProductPostgresqlCreate)
	exampleProductUsecaseGetByID := exampleProductUseCase.NewExampleProductUsecaseGetByID(exampleProductPostgresqlGetByID)
	exampleProductUsecaseGetPaginated := exampleProductUseCase.NewExampleProductUsecaseGetPaginated(exampleProductPostgresqlGetPaginated)
	exampleProductUsecaseUpdate := exampleProductUseCase.NewExampleProductUsecaseUpdate(exampleProductPostgresqlUpdate, exampleProductPostgresqlGetByID)
	exampleProductUsecaseDelete := exampleProductUseCase.NewExampleProductUsecaseDelete(exampleProductPostgresqlDelete, exampleProductPostgresqlGetByID)

	// Adapters In - Fiber
	authMiddleware := sharedFiber.NewAuth()

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
