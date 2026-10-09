package main

import (
	_ "github.com/kelsonwinith/learn.go-hexagonal-architecture/docs"
	bootstrap "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/bootstrap"
)

// ============================================================================
// Functions
// ============================================================================

// @title Hexagonal Architecture Go API
// @version 1.0
// @description This is a sample server following Hexagonal Architecture.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the JWT token.
func main() {
	// Config
	config := bootstrap.InitConfig()

	// Database
	db := bootstrap.InitDatabase(config)
	defer bootstrap.CloseDatabase(db)

	// HTTP App
	app := bootstrap.InitApp()

	// Middleware
	bootstrap.InitMiddleware(app)

	// Auth (JWT token service + bearer middleware)
	tokenService, authMiddleware := bootstrap.InitAuth(config)

	// Modules
	bootstrap.InitModules(app, db, tokenService, authMiddleware)

	// Server
	bootstrap.RunServer(app, config)
}
