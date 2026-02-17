package main

import "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/bootstrap"

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
func main() {
	bootstrap.Run()
}
