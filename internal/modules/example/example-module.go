package example

import (
	fiber "github.com/gofiber/fiber/v2"
	exampleFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/in/fiber"
	examplePostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/out/postgresql"
	exampleUseCase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/application"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

func Init(app *fiber.App, db *gorm.DB) {
	// Adapters Out - PostgreSQL
	postgresql := sharedPostgresql.NewPostgresql(db)
	postgresqlTransaction := sharedPostgresql.NewPostgresqlTransaction(postgresql)

	examplePostgresqlCreate := examplePostgresql.NewExamplePostgresqlCreate(postgresql)
	examplePostgresqlGetAll := examplePostgresql.NewExamplePostgresqlGetAll(postgresql)
	examplePostgresqlGetByID := examplePostgresql.NewExamplePostgresqlGetByID(postgresql)
	examplePostgresqlUpdate := examplePostgresql.NewExamplePostgresqlUpdate(postgresql)
	examplePostgresqlDelete := examplePostgresql.NewExamplePostgresqlDelete(postgresql)
	examplePostgresqlCreateMultiple := examplePostgresql.NewExamplePostgresqlCreateMultiple(postgresql)

	// Use Cases
	exampleUsecaseCreate := exampleUseCase.NewExampleUsecaseCreate(examplePostgresqlCreate)
	exampleUsecaseGetAll := exampleUseCase.NewExampleUsecaseGetAll(examplePostgresqlGetAll)
	exampleUsecaseGetByID := exampleUseCase.NewExampleUsecaseGetByID(examplePostgresqlGetByID)
	exampleUsecaseUpdate := exampleUseCase.NewExampleUsecaseUpdate(examplePostgresqlUpdate, examplePostgresqlGetByID)
	exampleUsecaseDelete := exampleUseCase.NewExampleUsecaseDelete(examplePostgresqlDelete)
	exampleUsecaseCreateMultiple := exampleUseCase.NewExampleUsecaseCreateMultiple(postgresqlTransaction, examplePostgresqlCreateMultiple)

	// Adapters In - Fiber
	exampleFiberCreate := exampleFiber.NewExampleFiberCreate(exampleUsecaseCreate)
	exampleFiberGetAll := exampleFiber.NewExampleFiberGetAll(exampleUsecaseGetAll)
	exampleFiberGetByID := exampleFiber.NewExampleFiberGetByID(exampleUsecaseGetByID)
	exampleFiberUpdate := exampleFiber.NewExampleFiberUpdate(exampleUsecaseUpdate)
	exampleFiberDelete := exampleFiber.NewExampleFiberDelete(exampleUsecaseDelete)
	exampleFiberCreateMultiple := exampleFiber.NewExampleFiberCreateMultiple(exampleUsecaseCreateMultiple)

	routes := app.Group("/api/v1/example")
	routes.Post("/", exampleFiberCreate.Handle)
	routes.Post("/batch", exampleFiberCreateMultiple.Handle)
	routes.Get("/", exampleFiberGetAll.Handle)
	routes.Get("/:id", exampleFiberGetByID.Handle)
	routes.Put("/:id", exampleFiberUpdate.Handle)
	routes.Delete("/:id", exampleFiberDelete.Handle)
}
