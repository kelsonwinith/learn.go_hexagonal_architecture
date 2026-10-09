package bootstrap

import (
	fiber "github.com/gofiber/fiber/v3"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Functions
// ============================================================================

func InitApp() *fiber.App {
	return fiber.New(fiber.Config{
		StructValidator: sharedFiber.NewValidator(),
	})
}
