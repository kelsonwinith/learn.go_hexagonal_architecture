package fiber

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	fiber "github.com/gofiber/fiber/v3"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

func TestAuthDefaultSuccess(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth())
	app.Get("/protected", func(c fiber.Ctx) error {
		userID := GetAuthUserID(c)
		token := GetAuthToken(c)
		return ResponseSuccess(c, map[string]any{
			"user_id": userID,
			"token":   token,
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("example-user-id", "42")

	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	assertStatus(t, res, http.StatusOK)
	assertResponseBody(t, res, true, map[string]any{
		"user_id": float64(42),
		"token":   "42",
	}, nil)
}

func TestAuthMissingHeader(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth())
	app.Get("/protected", func(c fiber.Ctx) error {
		return ResponseSuccess(c, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)

	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	assertStatus(t, res, http.StatusUnauthorized)

	var body responseBase
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if body.Error == nil || body.Error.ID != sharedDomain.FiberErrMissingAuthHeader.ID {
		t.Fatalf("Error.ID = %v, want %s", body.Error, sharedDomain.FiberErrMissingAuthHeader.ID)
	}
	if body.Error.Type != sharedDomain.ErrorTypeUnauthorized {
		t.Fatalf("Error.Type = %v, want %s", body.Error.Type, sharedDomain.ErrorTypeUnauthorized)
	}
}

func TestAuthInvalidUserID(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth())
	app.Get("/protected", func(c fiber.Ctx) error {
		return ResponseSuccess(c, "ok")
	})

	for _, invalidID := range []string{"not-a-number", "0", "-5", "abc"} {
		t.Run(invalidID, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			req.Header.Set("example-user-id", invalidID)

			res, err := app.Test(req)
			if err != nil {
				t.Fatalf("app.Test() error = %v", err)
			}
			defer res.Body.Close()

			assertStatus(t, res, http.StatusUnauthorized)

			var body responseBase
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				t.Fatalf("Decode() error = %v", err)
			}

			if body.Error == nil || body.Error.ID != sharedDomain.FiberErrInvalidAuthHeader.ID {
				t.Fatalf("Error.ID = %v, want %s", body.Error, sharedDomain.FiberErrInvalidAuthHeader.ID)
			}
		})
	}
}

func TestAuthCustomScheme(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth(AuthConfig{
		Header: "Authorization",
		Scheme: "Bearer",
	}))
	app.Get("/protected", func(c fiber.Ctx) error {
		return ResponseSuccess(c, "ok")
	})

	// Invalid scheme
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Basic invalid-scheme-token")

	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	assertStatus(t, res, http.StatusUnauthorized)

	// Empty token after scheme
	reqEmpty := httptest.NewRequest(http.MethodGet, "/protected", nil)
	reqEmpty.Header.Set("Authorization", "Bearer ")

	resEmpty, err := app.Test(reqEmpty)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer resEmpty.Body.Close()

	assertStatus(t, resEmpty, http.StatusUnauthorized)
}

func TestAuthCustomHeader(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth(AuthConfig{
		Header: "X-API-Key",
	}))
	app.Get("/protected", func(c fiber.Ctx) error {
		token := GetAuthToken(c)
		return ResponseSuccess(c, map[string]string{"key": token})
	})

	// Valid API key
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("X-API-Key", "api-secret-999")

	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	assertStatus(t, res, http.StatusOK)
	assertResponseBody(t, res, true, map[string]any{"key": "api-secret-999"}, nil)

	// Missing API key
	reqMissing := httptest.NewRequest(http.MethodGet, "/protected", nil)
	resMissing, err := app.Test(reqMissing)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer resMissing.Body.Close()

	assertStatus(t, resMissing, http.StatusUnauthorized)
}

func TestAuthCustomValidator(t *testing.T) {
	customErr := sharedDomain.NewError(sharedDomain.Unauthorized, "AUTH001", "token expired", nil)

	app := fiber.New()
	app.Use(NewAuth(AuthConfig{
		Header: "Authorization",
		Scheme: "Bearer",
		Validator: func(c fiber.Ctx, token string) (bool, error) {
			if token == "expired-token" {
				return false, customErr
			}
			if token == "secret-123" {
				return true, nil
			}
			return false, nil
		},
	}))
	app.Get("/protected", func(c fiber.Ctx) error {
		return ResponseSuccess(c, "access granted")
	})

	// Valid token
	reqValid := httptest.NewRequest(http.MethodGet, "/protected", nil)
	reqValid.Header.Set("Authorization", "Bearer secret-123")
	resValid, err := app.Test(reqValid)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer resValid.Body.Close()
	assertStatus(t, resValid, http.StatusOK)

	// Invalid token (returns FiberErrUnauthorized)
	reqInvalid := httptest.NewRequest(http.MethodGet, "/protected", nil)
	reqInvalid.Header.Set("Authorization", "Bearer wrong-token")
	resInvalid, err := app.Test(reqInvalid)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer resInvalid.Body.Close()
	assertStatus(t, resInvalid, http.StatusUnauthorized)

	var bodyInvalid responseBase
	if err := json.NewDecoder(resInvalid.Body).Decode(&bodyInvalid); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if bodyInvalid.Error.ID != sharedDomain.FiberErrUnauthorized.ID {
		t.Fatalf("Error.ID = %s, want %s", bodyInvalid.Error.ID, sharedDomain.FiberErrUnauthorized.ID)
	}

	// Token with custom error
	reqExpired := httptest.NewRequest(http.MethodGet, "/protected", nil)
	reqExpired.Header.Set("Authorization", "Bearer expired-token")
	resExpired, err := app.Test(reqExpired)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer resExpired.Body.Close()
	assertStatus(t, resExpired, http.StatusUnauthorized)

	var bodyExpired responseBase
	if err := json.NewDecoder(resExpired.Body).Decode(&bodyExpired); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if bodyExpired.Error.ID != "AUTH001" {
		t.Fatalf("Error.ID = %s, want AUTH001", bodyExpired.Error.ID)
	}
}

func TestAuthNextSkip(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth(AuthConfig{
		Next: func(c fiber.Ctx) bool {
			return c.Path() == "/public"
		},
	}))
	app.Get("/public", func(c fiber.Ctx) error {
		return ResponseSuccess(c, "public content")
	})

	req := httptest.NewRequest(http.MethodGet, "/public", nil)
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	assertStatus(t, res, http.StatusOK)
}

func TestGetAuthToken(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth(AuthConfig{
		Header:     "Authorization",
		Scheme:     "Bearer",
		ContextKey: "custom_key",
	}))
	app.Get("/test", func(c fiber.Ctx) error {
		token := GetAuthToken(c, "custom_key")
		emptyToken := GetAuthToken(c, "non_existent_key")
		return ResponseSuccess(c, map[string]string{
			"custom": token,
			"empty":  emptyToken,
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer test-val")

	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	assertStatus(t, res, http.StatusOK)
	assertResponseBody(t, res, true, map[string]any{
		"custom": "test-val",
		"empty":  "",
	}, nil)
}

func TestAuthCustomErrorHandler(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth(AuthConfig{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.Status(fiber.StatusForbidden).SendString("forbidden: " + err.Error())
		},
	}))
	app.Get("/protected", func(c fiber.Ctx) error {
		return ResponseSuccess(c, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	assertStatus(t, res, http.StatusForbidden)
}

func TestAuthUnregisteredErrorFallback(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuth(AuthConfig{
		Header: "Authorization",
		Scheme: "Bearer",
		Validator: func(c fiber.Ctx, token string) (bool, error) {
			return false, errors.New("raw unhandled error")
		},
	}))
	app.Get("/protected", func(c fiber.Ctx) error {
		return ResponseSuccess(c, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer some-token")
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	assertStatus(t, res, http.StatusInternalServerError)
}
