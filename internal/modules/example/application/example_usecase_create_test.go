package application

import (
	"context"
	"errors"
	"testing"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type fakePostgresCreate struct {
	err error
}

func (f *fakePostgresCreate) Execute(ctx context.Context, example *exampleDomain.Example) error {
	if f.err != nil {
		return f.err
	}
	example.ID = "generated-id"
	return nil
}

func TestExampleUsecaseCreate_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		repo := &fakePostgresCreate{}
		uc := NewExampleUsecaseCreate(repo)

		input := exampleDomain.Example{
			Name:        "John Doe",
			Description: "A valid description",
			CreatedBy:   100,
		}

		res, err := uc.Execute(ctx, input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res == nil {
			t.Fatal("expected result, got nil")
		}
		if res.Name != "John Doe" {
			t.Fatalf("expected Name 'John Doe', got %s", res.Name)
		}
		if res.CreatedBy != 100 || res.UpdatedBy != 100 {
			t.Fatalf("expected CreatedBy/UpdatedBy 100, got %d/%d", res.CreatedBy, res.UpdatedBy)
		}
	})

	t.Run("domain validation failure", func(t *testing.T) {
		repo := &fakePostgresCreate{}
		uc := NewExampleUsecaseCreate(repo)

		input := exampleDomain.Example{
			Name: "SingleNameOnly",
		}

		_, err := uc.Execute(ctx, input)
		if !errors.Is(err, exampleDomain.ExampleErrInvalidName) {
			t.Fatalf("expected ExampleErrInvalidName, got %v", err)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		repoErr := errors.New("db connection failed")
		repo := &fakePostgresCreate{err: repoErr}
		uc := NewExampleUsecaseCreate(repo)

		input := exampleDomain.Example{
			Name:      "John Doe",
			CreatedBy: 1,
		}

		_, err := uc.Execute(ctx, input)
		if !errors.Is(err, repoErr) {
			t.Fatalf("expected repo error, got %v", err)
		}
	})
}
