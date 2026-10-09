package exampleUser

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleUserFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/in/fiber"
	exampleUserPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/out/postgresql"
	exampleUserUseCase "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/application"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedPostgresql "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/postgresql"
	gorm "gorm.io/gorm"
)

// ============================================================================
// Functions
// ============================================================================

func Init(app *fiber.App, db *gorm.DB) exampleUserDomain.ExampleUserUsecaseGetByID {
	// Adapters Out - PostgreSQL
	postgresql := sharedPostgresql.NewPostgresql(db)

	exampleUserPostgresqlCreate := exampleUserPostgresql.NewExampleUserPostgresqlCreate(postgresql)
	exampleUserPostgresqlGetByID := exampleUserPostgresql.NewExampleUserPostgresqlGetByID(postgresql)
	exampleUserPostgresqlGetByEmail := exampleUserPostgresql.NewExampleUserPostgresqlGetByEmail(postgresql)

	// Use Cases
	exampleUserUsecaseRegister := exampleUserUseCase.NewExampleUserUsecaseRegister(exampleUserPostgresqlCreate)
	exampleUserUsecaseLogin := exampleUserUseCase.NewExampleUserUsecaseLogin(exampleUserPostgresqlGetByEmail)
	exampleUserUsecaseGetByID := exampleUserUseCase.NewExampleUserUsecaseGetByID(exampleUserPostgresqlGetByID)

	// Adapters In - Fiber
	exampleUserFiberRegister := exampleUserFiber.NewExampleUserFiberRegister(exampleUserUsecaseRegister)
	exampleUserFiberLogin := exampleUserFiber.NewExampleUserFiberLogin(exampleUserUsecaseLogin)
	exampleUserFiberGetByID := exampleUserFiber.NewExampleUserFiberGetByID(exampleUserUsecaseGetByID)

	exampleUserRoutes := app.Group("/api/v1/exampleuser")
	exampleUserRoutes.Post("/register", exampleUserFiberRegister.Handle)
	exampleUserRoutes.Post("/login", exampleUserFiberLogin.Handle)
	exampleUserRoutes.Get("/:id", exampleUserFiberGetByID.Handle)

	return exampleUserUsecaseGetByID
}
