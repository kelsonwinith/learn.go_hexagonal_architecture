package application

import (
	"context"
	"errors"
	"testing"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type fakePostgresDelete struct {
	deletedID string
	deletedBy int64
	err       error
}

func (f *fakePostgresDelete) Execute(ctx context.Context, id string, deletedBy int64) error {
	f.deletedID = id
	f.deletedBy = deletedBy
	return f.err
}

func TestExampleUsecaseDelete_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("success when deleted by creator", func(t *testing.T) {
		existing := &exampleDomain.Example{
			ID:        "123",
			Name:      "To Delete",
			CreatedBy: 42,
		}

		getByIDRepo := &fakePostgresGetByID{result: existing}
		deleteRepo := &fakePostgresDelete{}
		uc := NewExampleUsecaseDelete(deleteRepo, getByIDRepo)

		err := uc.Execute(ctx, "123", 42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if deleteRepo.deletedID != "123" {
			t.Fatalf("expected deletedID '123', got %s", deleteRepo.deletedID)
		}
		if deleteRepo.deletedBy != 42 {
			t.Fatalf("expected deletedBy 42, got %d", deleteRepo.deletedBy)
		}
	})

	t.Run("forbidden when deleted by non-creator", func(t *testing.T) {
		existing := &exampleDomain.Example{
			ID:        "123",
			Name:      "To Delete",
			CreatedBy: 42,
		}

		getByIDRepo := &fakePostgresGetByID{result: existing}
		deleteRepo := &fakePostgresDelete{}
		uc := NewExampleUsecaseDelete(deleteRepo, getByIDRepo)

		err := uc.Execute(ctx, "123", 99)
		if !errors.Is(err, exampleDomain.ExampleErrForbidden) {
			t.Fatalf("expected ExampleErrForbidden, got %v", err)
		}
	})

	t.Run("error when record not found", func(t *testing.T) {
		getByIDRepo := &fakePostgresGetByID{err: exampleDomain.ExampleErrNotFound}
		deleteRepo := &fakePostgresDelete{}
		uc := NewExampleUsecaseDelete(deleteRepo, getByIDRepo)

		err := uc.Execute(ctx, "non-existent", 42)
		if !errors.Is(err, exampleDomain.ExampleErrNotFound) {
			t.Fatalf("expected ExampleErrNotFound, got %v", err)
		}
	})

	t.Run("repository delete error", func(t *testing.T) {
		existing := &exampleDomain.Example{
			ID:        "123",
			Name:      "To Delete",
			CreatedBy: 42,
		}

		repoErr := errors.New("database delete failed")
		getByIDRepo := &fakePostgresGetByID{result: existing}
		deleteRepo := &fakePostgresDelete{err: repoErr}
		uc := NewExampleUsecaseDelete(deleteRepo, getByIDRepo)

		err := uc.Execute(ctx, "123", 42)
		if !errors.Is(err, repoErr) {
			t.Fatalf("expected repo error, got %v", err)
		}
	})
}
