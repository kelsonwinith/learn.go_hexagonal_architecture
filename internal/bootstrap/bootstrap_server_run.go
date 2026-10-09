package bootstrap

import (
	errors "errors"
	log "log"
	http "net/http"
	os "os"
	signal "os/signal"
	syscall "syscall"
	time "time"

	fiber "github.com/gofiber/fiber/v3"
	config "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
)

// ============================================================================
// Functions
// ============================================================================

func RunServer(app *fiber.App, cfg *config.Config) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on port %s", cfg.App.Port)
		if err := app.Listen(":" + cfg.App.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server gracefully...")

	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited cleanly")
}
