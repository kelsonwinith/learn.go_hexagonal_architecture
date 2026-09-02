package application

import (
	"context"
	"errors"
	"testing"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func TestExampleUsecaseGetByIDExecute(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		expected := &exampleDomain.Example{
			ID:   "123",
			Name: "Test Name",
		}
		repo := &fakePostgresGetByID{result: expected}
		uc := NewExampleUsecaseGetByID(repo)

		res, err := uc.Execute(ctx, "123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.ID != "123" {
			t.Fatalf("expected ID 123, got %s", res.ID)
		}
	})

	t.Run("not found", func(t *testing.T) {
		repo := &fakePostgresGetByID{err: exampleDomain.ExampleErrNotFound}
		uc := NewExampleUsecaseGetByID(repo)

		_, err := uc.Execute(ctx, "404")
		if !errors.Is(err, exampleDomain.ExampleErrNotFound) {
			t.Fatalf("expected ExampleErrNotFound, got %v", err)
		}
	})
}
