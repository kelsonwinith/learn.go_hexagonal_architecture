package fiber

import (
	strings "strings"

	fiber "github.com/gofiber/fiber/v3"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Constants
// ============================================================================

const (
	DefaultAuthHeader     = "Authorization"
	DefaultAuthContextKey = "auth_user_id"
	DefaultAuthScheme     = "Bearer"
)

// ============================================================================
// Types
// ============================================================================

// AuthConfig defines configuration options for the Auth middleware.
type AuthConfig struct {
	// Next defines a function to skip this middleware when returning true.
	Next func(c fiber.Ctx) bool

	// TokenService verifies the bearer token and returns its claims.
	TokenService sharedDomain.TokenService

	// Header is the HTTP request header carrying the credentials.
	// Defaults to "Authorization".
	Header string

	// Scheme is the expected authentication scheme (e.g. "Bearer").
	// Defaults to "Bearer".
	Scheme string

	// ContextKey is the key used to store the authenticated subject in c.Locals.
	// Defaults to "auth_user_id".
	ContextKey string

	// ErrorHandler is called when authentication fails.
	// Defaults to ResponseError(c, err).
	ErrorHandler func(c fiber.Ctx, err error) error
}

// ============================================================================
// Constructors
// ============================================================================

// NewAuth creates a Fiber middleware handler that validates a bearer token.
func NewAuth(config AuthConfig) fiber.Handler {
	if config.Header == "" {
		config.Header = DefaultAuthHeader
	}
	if config.Scheme == "" {
		config.Scheme = DefaultAuthScheme
	}
	if config.ContextKey == "" {
		config.ContextKey = DefaultAuthContextKey
	}

	errorHandler := config.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c fiber.Ctx, err error) error {
			return ResponseError(c, err)
		}
	}

	return func(c fiber.Ctx) error {
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		header := strings.TrimSpace(c.Get(config.Header))
		if header == "" {
			return errorHandler(c, sharedDomain.FiberErrMissingAuthHeader)
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], config.Scheme) || strings.TrimSpace(parts[1]) == "" {
			return errorHandler(c, sharedDomain.FiberErrInvalidAuthHeader)
		}

		claims, err := config.TokenService.Parse(c.Context(), strings.TrimSpace(parts[1]))
		if err != nil {
			return errorHandler(c, sharedDomain.FiberErrUnauthorized)
		}

		c.Locals(config.ContextKey, claims.Subject)

		return c.Next()
	}
}

// ============================================================================
// Functions
// ============================================================================

// GetAuthUserID retrieves the authenticated subject from Fiber context locals.
func GetAuthUserID(c fiber.Ctx, contextKey ...string) string {
	key := DefaultAuthContextKey
	if len(contextKey) > 0 && contextKey[0] != "" {
		key = contextKey[0]
	}

	if subject, ok := c.Locals(key).(string); ok {
		return subject
	}

	return ""
}
