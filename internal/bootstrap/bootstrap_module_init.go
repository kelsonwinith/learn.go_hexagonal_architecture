package bootstrap

import (
	fiber "github.com/gofiber/fiber/v3"
	gorm "gorm.io/gorm"

	exampleOrder "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder"
	exampleOrderExampleProduct "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/exampleproduct"
	exampleOrderExampleUser "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/out/exampleuser"
	exampleProduct "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct"
	exampleUser "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser"
)

// ============================================================================
// Functions
// ============================================================================

func InitModules(app *fiber.App, db *gorm.DB) {
	// Services exposed by other modules and consumed by exampleOrder
	exampleUserService := exampleUser.Init(app, db)
	exampleProductService := exampleProduct.Init(app, db)

	// Cross-module adapters: exampleOrder consumes the exampleUser and exampleProduct services
	exampleOrderUserReader := exampleOrderExampleUser.NewExampleOrderUserReader(exampleUserService)
	exampleOrderProductReader := exampleOrderExampleProduct.NewExampleOrderProductReader(exampleProductService)

	exampleOrder.Init(app, db, exampleOrderUserReader, exampleOrderProductReader)
}
