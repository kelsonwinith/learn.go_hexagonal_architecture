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

func TestResponseSuccess(t *testing.T) {
	app := fiber.New()
	app.Get("/success", func(c fiber.Ctx) error {
		return ResponseSuccess(c, map[string]string{"id": "1"})
	})

	res := performRequest(t, app, http.MethodGet, "/success")
	defer res.Body.Close()

	assertStatus(t, res, http.StatusOK)
	assertResponseBody(t, res, true, map[string]any{"id": "1"}, nil)
}

func TestResponseCreated(t *testing.T) {
	app := fiber.New()
	app.Post("/created", func(c fiber.Ctx) error {
		return ResponseCreated(c, map[string]string{"id": "1"})
	})

	res := performRequest(t, app, http.MethodPost, "/created")
	defer res.Body.Close()

	assertStatus(t, res, http.StatusCreated)
	assertResponseBody(t, res, true, map[string]any{"id": "1"}, nil)
}

func TestResponseNoContent(t *testing.T) {
	app := fiber.New()
	app.Delete("/empty", func(c fiber.Ctx) error {
		return ResponseNoContent(c)
	})

	res := performRequest(t, app, http.MethodDelete, "/empty")
	defer res.Body.Close()

	assertStatus(t, res, http.StatusNoContent)
}

func TestResponseError(t *testing.T) {
	appErr := sharedDomain.NewError(sharedDomain.BadRequest, "E001", "bad request", nil)
	app := fiber.New()
	app.Get("/error", func(c fiber.Ctx) error {
		return ResponseError(c, appErr)
	})

	res := performRequest(t, app, http.MethodGet, "/error")
	defer res.Body.Close()

	assertStatus(t, res, http.StatusBadRequest)
	assertResponseBody(t, res, false, nil, &responseBaseErrorBody{
		Type:    sharedDomain.ErrorTypeBadRequest,
		ID:      "E001",
		Message: "bad request",
	})
}

func TestResponseErrorFallback(t *testing.T) {
	app := fiber.New()
	app.Get("/error", func(c fiber.Ctx) error {
		return ResponseError(c, errors.New("unexpected"))
	})

	res := performRequest(t, app, http.MethodGet, "/error")
	defer res.Body.Close()

	assertStatus(t, res, http.StatusInternalServerError)
	assertResponseBody(t, res, false, nil, &responseBaseErrorBody{
		Type:    sharedDomain.ErrorTypeInternalServer,
		ID:      sharedDomain.SystemErrInternal.ID,
		Message: sharedDomain.SystemErrInternal.Message,
	})
}

func performRequest(t *testing.T, app *fiber.App, method, target string) *http.Response {
	t.Helper()

	res, err := app.Test(httptest.NewRequest(method, target, nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	return res
}

func assertStatus(t *testing.T, res *http.Response, want int) {
	t.Helper()

	if res.StatusCode != want {
		t.Fatalf("StatusCode = %d, want %d", res.StatusCode, want)
	}
}

func assertResponseBody(t *testing.T, res *http.Response, wantSuccess bool, wantData, wantError any) {
	t.Helper()

	var body responseBase
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if body.Success != wantSuccess {
		t.Fatalf("Success = %t, want %t", body.Success, wantSuccess)
	}

	assertJSONValue(t, "Data", body.Data, wantData)
	assertJSONValue(t, "Error", body.Error, wantError)
}

func assertJSONValue(t *testing.T, name string, got, want any) {
	t.Helper()

	gotJSON, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("Marshal(%s got) error = %v", name, err)
	}

	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("Marshal(%s want) error = %v", name, err)
	}

	if string(gotJSON) != string(wantJSON) {
		t.Fatalf("%s = %s, want %s", name, gotJSON, wantJSON)
	}
}
