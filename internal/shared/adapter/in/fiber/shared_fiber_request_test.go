package fiber

import (
	http "net/http"
	httptest "net/http/httptest"
	strings "strings"
	testing "testing"

	fiber "github.com/gofiber/fiber/v3"
)

func TestBind(t *testing.T) {
	type uri struct {
		ID string `uri:"id" validate:"required,uuid4"`
	}
	type query struct {
		Page int `query:"page" validate:"required,min=1"`
	}
	type body struct {
		Name string `json:"name" validate:"required"`
	}

	app := fiber.New(fiber.Config{StructValidator: NewValidator()})
	app.Put("/users/:id", func(c fiber.Ctx) error {
		_, err := Bind[uri, query, body](c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

	tests := []struct {
		name       string
		target     string
		body       string
		wantStatus int
	}{
		{
			name:       "valid params query and body",
			target:     "/users/2f9b6454-bbda-4c98-a709-732bb2163356?page=1",
			body:       `{"name":"Jane Doe"}`,
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "invalid params",
			target:     "/users/not-a-uuid?page=1",
			body:       `{"name":"Jane Doe"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid query",
			target:     "/users/2f9b6454-bbda-4c98-a709-732bb2163356?page=0",
			body:       `{"name":"Jane Doe"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid body",
			target:     "/users/2f9b6454-bbda-4c98-a709-732bb2163356?page=1",
			body:       `{}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, tt.target, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")

			res, err := app.Test(req)
			if err != nil {
				t.Fatalf("app.Test() error = %v", err)
			}
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Fatalf("StatusCode = %d, want %d", res.StatusCode, tt.wantStatus)
			}
		})
	}
}
