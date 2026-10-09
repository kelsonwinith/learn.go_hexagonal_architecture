package bootstrap

import (
	swaggo "github.com/gofiber/contrib/v3/swaggo"
	fiber "github.com/gofiber/fiber/v3"
	cors "github.com/gofiber/fiber/v3/middleware/cors"
	logger "github.com/gofiber/fiber/v3/middleware/logger"
)

// ============================================================================
// Functions
// ============================================================================

func InitMiddleware(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	// API documentation: Swagger UI, tags and operations collapsed by default
	swaggerConfig := swaggo.ConfigDefault
	swaggerConfig.DocExpansion = "none"
	swaggerConfig.DeepLinking = false
	swaggerConfig.PersistAuthorization = true
	app.Get("/swagger/*", swaggo.New(swaggerConfig))
}
