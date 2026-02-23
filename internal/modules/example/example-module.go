package example

import (
	fiber "github.com/gofiber/fiber/v2"
	sqlx "github.com/jmoiron/sqlx"
	exampleFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapters/in/fiber"
	examplePostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapters/out/postgresql"
	exampleUseCase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/application"
	postgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

func Init(app *fiber.App, db *sqlx.DB) {
	// Adapters Out - PostgreSQL
	postgresql := postgresql.NewPostgresql(db)
	createPostgresql := examplePostgresql.NewExampleCreate(postgresql)
	getAllPostgresql := examplePostgresql.NewExampleGetAll(postgresql)
	getByIDPostgresql := examplePostgresql.NewExampleGetByID(postgresql)
	updatePostgresql := examplePostgresql.NewExampleUpdate(postgresql)
	deletePostgresql := examplePostgresql.NewExampleDelete(postgresql)

	// Use Cases
	createUseCase := exampleUseCase.NewCreateExampleUseCase(createPostgresql)
	getAllUseCase := exampleUseCase.NewGetAllExamplesUseCase(getAllPostgresql)
	getByIDUseCase := exampleUseCase.NewGetExampleByIDUseCase(getByIDPostgresql)
	updateUseCase := exampleUseCase.NewUpdateExampleUseCase(updatePostgresql, getByIDPostgresql)
	deleteUseCase := exampleUseCase.NewDeleteExampleUseCase(deletePostgresql)

	// Adapters In - Fiber
	createHandler := exampleFiber.NewCreateExampleHandler(createUseCase)
	getAllHandler := exampleFiber.NewGetAllExamplesHandler(getAllUseCase)
	getByIDHandler := exampleFiber.NewGetExampleByIDHandler(getByIDUseCase)
	updateHandler := exampleFiber.NewUpdateExampleHandler(updateUseCase)
	deleteHandler := exampleFiber.NewDeleteExampleHandler(deleteUseCase)

	routes := app.Group("/example")
	routes.Post("/", createHandler.Handle)
	routes.Get("/", getAllHandler.Handle)
	routes.Get("/:id", getByIDHandler.Handle)
	routes.Put("/:id", updateHandler.Handle)
	routes.Delete("/:id", deleteHandler.Handle)
}
