package example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapters/postgresql"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/application"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/interfaces/http"
)

func Init(app *fiber.App, db *sqlx.DB) {
	// Adapters (Repository)
	repo := postgresql.NewExampleRepository(db)

	// Use Cases
	createUseCase := application.NewCreateExampleUseCase(repo)
	getAllUseCase := application.NewGetAllExamplesUseCase(repo)
	getByIDUseCase := application.NewGetExampleByIDUseCase(repo)
	updateUseCase := application.NewUpdateExampleUseCase(repo)
	deleteUseCase := application.NewDeleteExampleUseCase(repo)

	// Handlers
	createHandler := http.NewCreateExampleHandler(createUseCase)
	getAllHandler := http.NewGetAllExamplesHandler(getAllUseCase)
	getByIDHandler := http.NewGetExampleByIDHandler(getByIDUseCase)
	updateHandler := http.NewUpdateExampleHandler(updateUseCase)
	deleteHandler := http.NewDeleteExampleHandler(deleteUseCase)

	// Routes
	routes := app.Group("/examples")
	routes.Post("/", createHandler.Handle)
	routes.Get("/", getAllHandler.Handle)
	routes.Get("/:id", getByIDHandler.Handle)
	routes.Put("/:id", updateHandler.Handle)
	routes.Delete("/:id", deleteHandler.Handle)
}
