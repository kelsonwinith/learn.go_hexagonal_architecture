package fiber

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	fiber "github.com/gofiber/fiber/v3"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// Fake usecase implementations
type fakeCreateUsecase struct{}

func (f *fakeCreateUsecase) Execute(ctx context.Context, input exampleDomain.Example) (*exampleDomain.Example, error) {
	return &exampleDomain.Example{
		ID:          "123",
		Name:        input.Name,
		Description: input.Description,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

type fakeCreateMultipleUsecase struct{}

func (f *fakeCreateMultipleUsecase) Execute(ctx context.Context, examples []exampleDomain.Example) ([]*exampleDomain.Example, error) {
	res := make([]*exampleDomain.Example, len(examples))
	for i, ex := range examples {
		res[i] = &exampleDomain.Example{
			ID:          "123",
			Name:        ex.Name,
			Description: ex.Description,
			CreatedBy:   ex.CreatedBy,
			UpdatedBy:   ex.UpdatedBy,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}
	return res, nil
}

type fakeUpdateUsecase struct{}

func (f *fakeUpdateUsecase) Execute(ctx context.Context, input exampleDomain.Example) (*exampleDomain.Example, error) {
	if input.UpdatedBy != 123 && input.UpdatedBy != 99 {
		return nil, exampleDomain.ExampleErrForbidden
	}
	return &exampleDomain.Example{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		CreatedBy:   123,
		UpdatedBy:   input.UpdatedBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

type fakeDeleteUsecase struct{}

func (f *fakeDeleteUsecase) Execute(ctx context.Context, id string, userID int64) error {
	if userID != 123 && userID != 99 {
		return exampleDomain.ExampleErrForbidden
	}
	return nil
}

type fakeGetAllUsecase struct{}

func (f *fakeGetAllUsecase) Execute(ctx context.Context) ([]*exampleDomain.Example, error) {
	return []*exampleDomain.Example{}, nil
}

type fakeGetByIDUsecase struct{}

func (f *fakeGetByIDUsecase) Execute(ctx context.Context, id string) (*exampleDomain.Example, error) {
	return &exampleDomain.Example{
		ID:        id,
		Name:      "Test Example",
		CreatedBy: 123,
		UpdatedBy: 123,
	}, nil
}

func setupTestApp() *fiber.App {
	app := fiber.New(fiber.Config{
		StructValidator: sharedFiber.NewValidator(),
	})

	createHandler := NewExampleFiberCreate(&fakeCreateUsecase{})
	createMultipleHandler := NewExampleFiberCreateMultiple(&fakeCreateMultipleUsecase{})
	updateHandler := NewExampleFiberUpdate(&fakeUpdateUsecase{})
	deleteHandler := NewExampleFiberDelete(&fakeDeleteUsecase{})
	getAllHandler := NewExampleFiberGetAll(&fakeGetAllUsecase{})
	getByIDHandler := NewExampleFiberGetByID(&fakeGetByIDUsecase{})

	authMiddleware := sharedFiber.NewAuth()

	routes := app.Group("/api/v1/example")
	routes.Post("/", authMiddleware, createHandler.Handle)
	routes.Post("/batch", authMiddleware, createMultipleHandler.Handle)
	routes.Get("/", getAllHandler.Handle)
	routes.Get("/:id", getByIDHandler.Handle)
	routes.Put("/:id", authMiddleware, updateHandler.Handle)
	routes.Delete("/:id", authMiddleware, deleteHandler.Handle)

	return app
}

func TestExampleRoutes_AuthProtection(t *testing.T) {
	app := setupTestApp()

	tests := []struct {
		name        string
		method      string
		url         string
		body        string
		withAuth    bool
		userID      string
		wantStatus  int
	}{
		// Public routes (no auth needed)
		{
			name:       "GET /api/v1/example without auth succeeds",
			method:     http.MethodGet,
			url:        "/api/v1/example",
			withAuth:   false,
			wantStatus: http.StatusOK,
		},
		{
			name:       "GET /api/v1/example/:id without auth succeeds",
			method:     http.MethodGet,
			url:        "/api/v1/example/2f9b6454-bbda-4c98-a709-732bb2163356",
			withAuth:   false,
			wantStatus: http.StatusOK,
		},

		// Create route
		{
			name:       "POST /api/v1/example without auth is unauthorized",
			method:     http.MethodPost,
			url:        "/api/v1/example",
			body:       `{"name":"John Doe"}`,
			withAuth:   false,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "POST /api/v1/example with auth is authorized",
			method:     http.MethodPost,
			url:        "/api/v1/example",
			body:       `{"name":"John Doe"}`,
			withAuth:   true,
			userID:     "123",
			wantStatus: http.StatusCreated,
		},

		// CreateMultiple route
		{
			name:       "POST /api/v1/example/batch without auth is unauthorized",
			method:     http.MethodPost,
			url:        "/api/v1/example/batch",
			body:       `{"examples":[{"name":"John Doe"}]}`,
			withAuth:   false,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "POST /api/v1/example/batch with auth is authorized",
			method:     http.MethodPost,
			url:        "/api/v1/example/batch",
			body:       `{"examples":[{"name":"John Doe"}]}`,
			withAuth:   true,
			userID:     "123",
			wantStatus: http.StatusCreated,
		},

		// Update route
		{
			name:       "PUT /api/v1/example/:id without auth is unauthorized",
			method:     http.MethodPut,
			url:        "/api/v1/example/2f9b6454-bbda-4c98-a709-732bb2163356",
			body:       `{"name":"Jane Doe"}`,
			withAuth:   false,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "PUT /api/v1/example/:id with matching auth is authorized",
			method:     http.MethodPut,
			url:        "/api/v1/example/2f9b6454-bbda-4c98-a709-732bb2163356",
			body:       `{"name":"Jane Doe"}`,
			withAuth:   true,
			userID:     "123",
			wantStatus: http.StatusOK,
		},
		{
			name:       "PUT /api/v1/example/:id with non-matching auth is forbidden",
			method:     http.MethodPut,
			url:        "/api/v1/example/2f9b6454-bbda-4c98-a709-732bb2163356",
			body:       `{"name":"Jane Doe"}`,
			withAuth:   true,
			userID:     "999",
			wantStatus: http.StatusForbidden,
		},

		// Delete route
		{
			name:       "DELETE /api/v1/example/:id without auth is unauthorized",
			method:     http.MethodDelete,
			url:        "/api/v1/example/2f9b6454-bbda-4c98-a709-732bb2163356",
			withAuth:   false,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "DELETE /api/v1/example/:id with matching auth is authorized",
			method:     http.MethodDelete,
			url:        "/api/v1/example/2f9b6454-bbda-4c98-a709-732bb2163356",
			withAuth:   true,
			userID:     "123",
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "DELETE /api/v1/example/:id with non-matching auth is forbidden",
			method:     http.MethodDelete,
			url:        "/api/v1/example/2f9b6454-bbda-4c98-a709-732bb2163356",
			withAuth:   true,
			userID:     "999",
			wantStatus: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bodyReader *bytes.Reader
			if tt.body != "" {
				bodyReader = bytes.NewReader([]byte(tt.body))
			} else {
				bodyReader = bytes.NewReader([]byte{})
			}

			req := httptest.NewRequest(tt.method, tt.url, bodyReader)
			req.Header.Set("Content-Type", "application/json")
			if tt.withAuth {
				uid := tt.userID
				if uid == "" {
					uid = "123"
				}
				req.Header.Set("example-user-id", uid)
			}

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

func TestExampleCreate_UserIDRecorded(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/example", bytes.NewReader([]byte(`{"name":"Jane Doe"}`)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("example-user-id", "99")

	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("StatusCode = %d, want %d", res.StatusCode, http.StatusCreated)
	}

	type responseWrapper struct {
		Success bool `json:"success"`
		Data    struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			CreatedBy int64  `json:"created_by"`
			UpdatedBy int64  `json:"updated_by"`
		} `json:"data"`
	}

	var resData responseWrapper
	importJsonErr := json.NewDecoder(res.Body).Decode(&resData)
	if importJsonErr != nil {
		t.Fatalf("Decode() error = %v", importJsonErr)
	}

	if resData.Data.CreatedBy != 99 {
		t.Fatalf("CreatedBy = %d, want 99", resData.Data.CreatedBy)
	}
	if resData.Data.UpdatedBy != 99 {
		t.Fatalf("UpdatedBy = %d, want 99", resData.Data.UpdatedBy)
	}
}
