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

// @securityDefinitions.apikey UserIdAuth
// @in header
// @name example-user-id
// @description Pass the user ID as integer (e.g. 1).
func main() {
	bootstrap.Run()
}
