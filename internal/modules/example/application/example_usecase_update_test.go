package application

import (
	"context"
	"errors"
	"testing"
	"time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type fakePostgresGetByID struct {
	result *exampleDomain.Example
	err    error
}

func (f *fakePostgresGetByID) Execute(ctx context.Context, id string) (*exampleDomain.Example, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.result, nil
}

type fakePostgresUpdate struct {
	err error
}

func (f *fakePostgresUpdate) Execute(ctx context.Context, example *exampleDomain.Example) error {
	return f.err
}

func TestExampleUsecaseUpdateExecute(t *testing.T) {
	ctx := context.Background()

	t.Run("success when updated by creator", func(t *testing.T) {
		existing := &exampleDomain.Example{
			ID:          "123",
			Name:        "Old Name",
			Description: "Old Desc",
			CreatedBy:   42,
			UpdatedBy:   42,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		getByIDRepo := &fakePostgresGetByID{result: existing}
		updateRepo := &fakePostgresUpdate{}
		uc := NewExampleUsecaseUpdate(updateRepo, getByIDRepo)

		input := exampleDomain.Example{
			ID:          "123",
			Name:        "New Name",
			Description: "New Desc",
			UpdatedBy:   42,
		}

		res, err := uc.Execute(ctx, input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.Name != "New Name" {
			t.Fatalf("expected Name 'New Name', got %s", res.Name)
		}
		if res.UpdatedBy != 42 {
			t.Fatalf("expected UpdatedBy 42, got %d", res.UpdatedBy)
		}
	})

	t.Run("forbidden when updated by non-creator", func(t *testing.T) {
		existing := &exampleDomain.Example{
			ID:        "123",
			Name:      "Existing Name",
			CreatedBy: 42,
		}

		getByIDRepo := &fakePostgresGetByID{result: existing}
		updateRepo := &fakePostgresUpdate{}
		uc := NewExampleUsecaseUpdate(updateRepo, getByIDRepo)

		input := exampleDomain.Example{
			ID:        "123",
			Name:      "New Name",
			UpdatedBy: 99, // Different from CreatedBy
		}

		_, err := uc.Execute(ctx, input)
		if !errors.Is(err, exampleDomain.ExampleErrForbidden) {
			t.Fatalf("expected ExampleErrForbidden, got %v", err)
		}
	})

	t.Run("error when existing record not found", func(t *testing.T) {
		getByIDRepo := &fakePostgresGetByID{err: exampleDomain.ExampleErrNotFound}
		updateRepo := &fakePostgresUpdate{}
		uc := NewExampleUsecaseUpdate(updateRepo, getByIDRepo)

		input := exampleDomain.Example{
			ID:        "non-existent",
			Name:      "New Name",
			UpdatedBy: 42,
		}

		_, err := uc.Execute(ctx, input)
		if !errors.Is(err, exampleDomain.ExampleErrNotFound) {
			t.Fatalf("expected ExampleErrNotFound, got %v", err)
		}
	})

	t.Run("domain validation error", func(t *testing.T) {
		existing := &exampleDomain.Example{
			ID:        "123",
			Name:      "Valid Name",
			CreatedBy: 42,
		}

		getByIDRepo := &fakePostgresGetByID{result: existing}
		updateRepo := &fakePostgresUpdate{}
		uc := NewExampleUsecaseUpdate(updateRepo, getByIDRepo)

		input := exampleDomain.Example{
			ID:        "123",
			Name:      "SingleInvalidName",
			UpdatedBy: 42,
		}

		_, err := uc.Execute(ctx, input)
		if !errors.Is(err, exampleDomain.ExampleErrInvalidName) {
			t.Fatalf("expected ExampleErrInvalidName, got %v", err)
		}
	})

	t.Run("repository update error", func(t *testing.T) {
		existing := &exampleDomain.Example{
			ID:        "123",
			Name:      "Valid Name",
			CreatedBy: 42,
		}

		repoErr := errors.New("database update error")
		getByIDRepo := &fakePostgresGetByID{result: existing}
		updateRepo := &fakePostgresUpdate{err: repoErr}
		uc := NewExampleUsecaseUpdate(updateRepo, getByIDRepo)

		input := exampleDomain.Example{
			ID:        "123",
			Name:      "Updated Name",
			UpdatedBy: 42,
		}

		_, err := uc.Execute(ctx, input)
		if !errors.Is(err, repoErr) {
			t.Fatalf("expected repo error, got %v", err)
		}
	})
}
