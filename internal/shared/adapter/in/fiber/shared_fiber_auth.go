package fiber

import (
	"strconv"
	"strings"

	fiber "github.com/gofiber/fiber/v3"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Constants
// ============================================================================

const (
	DefaultAuthHeader     = "example-user-id"
	DefaultAuthContextKey = "auth_user_id"
	DefaultAuthScheme     = ""
)

// ============================================================================
// Variables
// ============================================================================
// DefaultAuthConfig is the default configuration for Auth middleware.
var DefaultAuthConfig = AuthConfig{
	Header:     DefaultAuthHeader,
	Scheme:     DefaultAuthScheme,
	ContextKey: DefaultAuthContextKey,
}

// ============================================================================
// Types
// ============================================================================
// AuthConfig defines configuration options for the Auth middleware.
type AuthConfig struct {
	// Next defines a function to skip this middleware when returning true.
	Next func(c fiber.Ctx) bool

	// Header is the HTTP request header key used to look up authentication credentials.
	// Defaults to "example-user-id".
	Header string

	// Scheme is the expected authentication scheme (e.g. "Bearer").
	// When Scheme is set, the header value must match "<Scheme> <token>".
	// When Scheme is empty (""), the full header value is treated as the token/id.
	// Defaults to "".
	Scheme string

	// ContextKey is the key used to store the extracted user ID / token in c.Locals.
	// Defaults to "auth_user_id".
	ContextKey string

	// Validator is an optional function to validate the token.
	// If nil and Header is "example-user-id", validates that the token is a positive integer.
	// If it returns (false, nil), FiberErrUnauthorized is returned.
	// If it returns (false, err), err is passed to ErrorHandler.
	Validator func(c fiber.Ctx, token string) (bool, error)

	// ErrorHandler is called when authentication fails.
	// Defaults to ResponseError(c, err).
	ErrorHandler func(c fiber.Ctx, err error) error
}

// ============================================================================
// Constructors
// ============================================================================
// NewAuth creates a Fiber middleware handler that performs authentication checks.
func NewAuth(config ...AuthConfig) fiber.Handler {
	cfg := DefaultAuthConfig
	if len(config) > 0 {
		userCfg := config[0]
		if userCfg.Next != nil {
			cfg.Next = userCfg.Next
		}
		if userCfg.Header != "" {
			cfg.Header = userCfg.Header
		}
		cfg.Scheme = userCfg.Scheme
		if userCfg.ContextKey != "" {
			cfg.ContextKey = userCfg.ContextKey
		}
		cfg.Validator = userCfg.Validator
		cfg.ErrorHandler = userCfg.ErrorHandler
	}

	errorHandler := cfg.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c fiber.Ctx, err error) error {
			return ResponseError(c, err)
		}
	}

	return func(c fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		authHeader := strings.TrimSpace(c.Get(cfg.Header))
		if authHeader == "" {
			return errorHandler(c, sharedDomain.FiberErrMissingAuthHeader)
		}

		token := authHeader
		if cfg.Scheme != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], cfg.Scheme) || strings.TrimSpace(parts[1]) == "" {
				return errorHandler(c, sharedDomain.FiberErrInvalidAuthHeader)
			}
			token = strings.TrimSpace(parts[1])
		}

		if cfg.Validator != nil {
			valid, err := cfg.Validator(c, token)
			if err != nil {
				return errorHandler(c, err)
			}
			if !valid {
				return errorHandler(c, sharedDomain.FiberErrUnauthorized)
			}
		} else if cfg.Header == DefaultAuthHeader {
			// By default, example-user-id must be a valid positive integer
			userID, err := strconv.ParseInt(token, 10, 64)
			if err != nil || userID <= 0 {
				return errorHandler(c, sharedDomain.FiberErrInvalidAuthHeader)
			}
			c.Locals(cfg.ContextKey, userID)
			c.Locals("auth_token", token)
			return c.Next()
		}

		// If token can be parsed as int64, also store as int64 in context
		if userID, err := strconv.ParseInt(token, 10, 64); err == nil {
			c.Locals(cfg.ContextKey, userID)
		} else {
			c.Locals(cfg.ContextKey, token)
		}
		c.Locals("auth_token", token)
		return c.Next()
	}
}

// ============================================================================
// Functions
// ============================================================================
// GetAuthUserID retrieves the authenticated user's integer ID from Fiber context locals.
func GetAuthUserID(c fiber.Ctx, contextKey ...string) int64 {
	key := DefaultAuthContextKey
	if len(contextKey) > 0 && contextKey[0] != "" {
		key = contextKey[0]
	}

	val := c.Locals(key)
	switch v := val.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case string:
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			return id
		}
	}
	return 0
}

// GetAuthToken retrieves the stored authentication token from Fiber context locals.
func GetAuthToken(c fiber.Ctx, contextKey ...string) string {
	key := "auth_token"
	if len(contextKey) > 0 && contextKey[0] != "" {
		key = contextKey[0]
	}

	val := c.Locals(key)
	if token, ok := val.(string); ok {
		return token
	}
	if id, ok := val.(int64); ok {
		return strconv.FormatInt(id, 10)
	}
	return ""
}
