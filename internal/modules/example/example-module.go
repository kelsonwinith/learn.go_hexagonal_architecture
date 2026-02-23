package example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapters/in/http"
	examplerepository "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapters/out/postgresql"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/application"
	sharedpostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapters/out/postgresql"
)

func Init(app *fiber.App, db *sqlx.DB) {
	// Repository
	postgresql := sharedpostgresql.NewPostgresql(db)
	createRepo := examplerepository.NewExampleCreate(postgresql)
	getAllRepo := examplerepository.NewExampleGetAll(postgresql)
	getByIDRepo := examplerepository.NewExampleGetByID(postgresql)
	updateRepo := examplerepository.NewExampleUpdate(postgresql)
	deleteRepo := examplerepository.NewExampleDelete(postgresql)

	// Use Cases
	createUseCase := application.NewCreateExampleUseCase(createRepo)
	getAllUseCase := application.NewGetAllExamplesUseCase(getAllRepo)
	getByIDUseCase := application.NewGetExampleByIDUseCase(getByIDRepo)
	updateUseCase := application.NewUpdateExampleUseCase(updateRepo, getByIDRepo)
	deleteUseCase := application.NewDeleteExampleUseCase(deleteRepo)

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
