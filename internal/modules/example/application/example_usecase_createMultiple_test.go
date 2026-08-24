package application

import (
	"context"
	"errors"
	"testing"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type fakePostgresTransaction struct {
	err error
}

func (f *fakePostgresTransaction) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if f.err != nil {
		return f.err
	}
	return fn(ctx)
}

type fakePostgresCreateMultiple struct {
	err error
}

func (f *fakePostgresCreateMultiple) Execute(ctx context.Context, examples []*exampleDomain.Example) error {
	if f.err != nil {
		return f.err
	}
	for i, ex := range examples {
		ex.ID = string(rune('1' + i))
	}
	return nil
}

func TestExampleUsecaseCreateMultiple_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("success creating multiple examples", func(t *testing.T) {
		tx := &fakePostgresTransaction{}
		repo := &fakePostgresCreateMultiple{}
		uc := NewExampleUsecaseCreateMultiple(tx, repo)

		inputs := []exampleDomain.Example{
			{Name: "Alice Smith", Description: "First description", CreatedBy: 10},
			{Name: "Bob Jones", Description: "Second description", CreatedBy: 10},
		}

		results, err := uc.Execute(ctx, inputs)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(results) != 2 {
			t.Fatalf("expected 2 results, got %d", len(results))
		}
		if results[0].Name != "Alice Smith" || results[0].CreatedBy != 10 {
			t.Fatalf("unexpected first result: %+v", results[0])
		}
		if results[1].Name != "Bob Jones" || results[1].CreatedBy != 10 {
			t.Fatalf("unexpected second result: %+v", results[1])
		}
	})

	t.Run("domain validation failure on one item", func(t *testing.T) {
		tx := &fakePostgresTransaction{}
		repo := &fakePostgresCreateMultiple{}
		uc := NewExampleUsecaseCreateMultiple(tx, repo)

		inputs := []exampleDomain.Example{
			{Name: "Alice Smith", Description: "Valid description", CreatedBy: 10},
			{Name: "SingleInvalidName", Description: "Invalid description", CreatedBy: 10},
		}

		_, err := uc.Execute(ctx, inputs)
		if !errors.Is(err, exampleDomain.ExampleErrInvalidName) {
			t.Fatalf("expected ExampleErrInvalidName, got %v", err)
		}
	})

	t.Run("transaction error", func(t *testing.T) {
		txErr := errors.New("transaction start failed")
		tx := &fakePostgresTransaction{err: txErr}
		repo := &fakePostgresCreateMultiple{}
		uc := NewExampleUsecaseCreateMultiple(tx, repo)

		inputs := []exampleDomain.Example{
			{Name: "Alice Smith", CreatedBy: 10},
		}

		_, err := uc.Execute(ctx, inputs)
		if !errors.Is(err, txErr) {
			t.Fatalf("expected transaction error, got %v", err)
		}
	})

	t.Run("repository create multiple error", func(t *testing.T) {
		repoErr := errors.New("db insert batch failed")
		tx := &fakePostgresTransaction{}
		repo := &fakePostgresCreateMultiple{err: repoErr}
		uc := NewExampleUsecaseCreateMultiple(tx, repo)

		inputs := []exampleDomain.Example{
			{Name: "Alice Smith", CreatedBy: 10},
		}

		_, err := uc.Execute(ctx, inputs)
		if !errors.Is(err, repoErr) {
			t.Fatalf("expected repository error, got %v", err)
		}
	})
}
