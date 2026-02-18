package example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapters/repository"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/application"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/interfaces/http"
)

func Init(app *fiber.App, db *sqlx.DB) {
	// Repository
	baseRepo := repository.NewExampleRepository(db)
	createRepo := repository.NewExampleCreateRepository(baseRepo)
	getAllRepo := repository.NewExampleGetAllRepository(baseRepo)
	getByIDRepo := repository.NewExampleGetByIDRepository(baseRepo)
	updateRepo := repository.NewExampleUpdateRepository(baseRepo)
	deleteRepo := repository.NewExampleDeleteRepository(baseRepo)

	// Use Cases
	var createUseCase domain.CreateExampleUseCase = application.NewCreateExampleUseCase(createRepo)
	var getAllUseCase domain.GetAllExamplesUseCase = application.NewGetAllExamplesUseCase(getAllRepo)
	var getByIDUseCase domain.GetExampleByIDUseCase = application.NewGetExampleByIDUseCase(getByIDRepo)
	var updateUseCase domain.UpdateExampleUseCase = application.NewUpdateExampleUseCase(updateRepo, getByIDRepo)
	var deleteUseCase domain.DeleteExampleUseCase = application.NewDeleteExampleUseCase(deleteRepo)

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
