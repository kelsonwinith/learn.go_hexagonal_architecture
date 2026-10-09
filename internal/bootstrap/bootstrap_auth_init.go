package bootstrap

import (
	fiber "github.com/gofiber/fiber/v3"
	config "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/config"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
	sharedJwt "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/out/jwt"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Functions
// ============================================================================

func InitAuth(cfg *config.Config) (sharedDomain.TokenService, fiber.Handler) {
	tokenService := sharedJwt.NewJwtService(cfg.JWT.Secret, cfg.JWT.ExpiryHours)
	authMiddleware := sharedFiber.NewAuth(sharedFiber.AuthConfig{TokenService: tokenService})

	return tokenService, authMiddleware
}
