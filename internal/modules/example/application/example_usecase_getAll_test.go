package application

import (
	"context"
	"errors"
	"testing"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type fakePostgresGetAll struct {
	results []*exampleDomain.Example
	err     error
}

func (f *fakePostgresGetAll) Execute(ctx context.Context) ([]*exampleDomain.Example, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.results, nil
}

func TestExampleUsecaseGetAllExecute(t *testing.T) {
	ctx := context.Background()

	t.Run("success with items", func(t *testing.T) {
		expected := []*exampleDomain.Example{
			{ID: "1", Name: "Item One"},
			{ID: "2", Name: "Item Two"},
		}
		repo := &fakePostgresGetAll{results: expected}
		uc := NewExampleUsecaseGetAll(repo)

		res, err := uc.Execute(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) != 2 {
			t.Fatalf("expected 2 items, got %d", len(res))
		}
	})

	t.Run("repository error", func(t *testing.T) {
		repoErr := errors.New("db error")
		repo := &fakePostgresGetAll{err: repoErr}
		uc := NewExampleUsecaseGetAll(repo)

		_, err := uc.Execute(ctx)
		if !errors.Is(err, repoErr) {
			t.Fatalf("expected repo error, got %v", err)
		}
	})
}
