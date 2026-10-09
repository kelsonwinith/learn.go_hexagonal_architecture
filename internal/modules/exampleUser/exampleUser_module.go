package exampleUser

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleUserFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/in/fiber"
	exampleUserPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/out/postgresql"
	exampleUserUseCase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/application"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Functions
// ============================================================================

func Init(app *fiber.App, db *gorm.DB) exampleUserDomain.ExampleUserUsecaseGetByID {
	// Adapters Out - PostgreSQL
	postgresql := sharedPostgresql.NewPostgresql(db)
	postgresqlTransaction := sharedPostgresql.NewPostgresqlTransaction(postgresql)

	exampleUserPostgresqlCreate := exampleUserPostgresql.NewExampleUserPostgresqlCreate(postgresql)
	exampleUserPostgresqlGetAll := exampleUserPostgresql.NewExampleUserPostgresqlGetAll(postgresql)
	exampleUserPostgresqlGetPaginated := exampleUserPostgresql.NewExampleUserPostgresqlGetPaginated(postgresql)
	exampleUserPostgresqlGetByID := exampleUserPostgresql.NewExampleUserPostgresqlGetByID(postgresql)
	exampleUserPostgresqlUpdate := exampleUserPostgresql.NewExampleUserPostgresqlUpdate(postgresql)
	exampleUserPostgresqlDelete := exampleUserPostgresql.NewExampleUserPostgresqlDelete(postgresql)
	exampleUserPostgresqlCreateMultiple := exampleUserPostgresql.NewExampleUserPostgresqlCreateMultiple(postgresql)

	// Use Cases
	exampleUserUsecaseCreate := exampleUserUseCase.NewExampleUserUsecaseCreate(exampleUserPostgresqlCreate)
	exampleUserUsecaseGetAll := exampleUserUseCase.NewExampleUserUsecaseGetAll(exampleUserPostgresqlGetAll)
	exampleUserUsecaseGetPaginated := exampleUserUseCase.NewExampleUserUsecaseGetPaginated(exampleUserPostgresqlGetPaginated)
	exampleUserUsecaseGetByID := exampleUserUseCase.NewExampleUserUsecaseGetByID(exampleUserPostgresqlGetByID)
	exampleUserUsecaseUpdate := exampleUserUseCase.NewExampleUserUsecaseUpdate(exampleUserPostgresqlUpdate, exampleUserPostgresqlGetByID)
	exampleUserUsecaseDelete := exampleUserUseCase.NewExampleUserUsecaseDelete(exampleUserPostgresqlDelete, exampleUserPostgresqlGetByID)
	exampleUserUsecaseCreateMultiple := exampleUserUseCase.NewExampleUserUsecaseCreateMultiple(postgresqlTransaction, exampleUserPostgresqlCreateMultiple)

	// Adapters In - Fiber
	authMiddleware := sharedFiber.NewAuth()

	exampleUserFiberCreate := exampleUserFiber.NewExampleUserFiberCreate(exampleUserUsecaseCreate)
	exampleUserFiberGetAll := exampleUserFiber.NewExampleUserFiberGetAll(exampleUserUsecaseGetAll)
	exampleUserFiberGetPaginated := exampleUserFiber.NewExampleUserFiberGetPaginated(exampleUserUsecaseGetPaginated)
	exampleUserFiberGetByID := exampleUserFiber.NewExampleUserFiberGetByID(exampleUserUsecaseGetByID)
	exampleUserFiberUpdate := exampleUserFiber.NewExampleUserFiberUpdate(exampleUserUsecaseUpdate)
	exampleUserFiberDelete := exampleUserFiber.NewExampleUserFiberDelete(exampleUserUsecaseDelete)
	exampleUserFiberCreateMultiple := exampleUserFiber.NewExampleUserFiberCreateMultiple(exampleUserUsecaseCreateMultiple)

	exampleUserRoutes := app.Group("/api/v1/exampleuser")
	exampleUserRoutes.Post("/", authMiddleware, exampleUserFiberCreate.Handle)
	exampleUserRoutes.Post("/batch", authMiddleware, exampleUserFiberCreateMultiple.Handle)
	exampleUserRoutes.Get("/", exampleUserFiberGetAll.Handle)
	exampleUserRoutes.Get("/paginated", exampleUserFiberGetPaginated.Handle)
	exampleUserRoutes.Get("/:id", exampleUserFiberGetByID.Handle)
	exampleUserRoutes.Put("/:id", authMiddleware, exampleUserFiberUpdate.Handle)
	exampleUserRoutes.Delete("/:id", authMiddleware, exampleUserFiberDelete.Handle)

	return exampleUserUsecaseGetByID
}
