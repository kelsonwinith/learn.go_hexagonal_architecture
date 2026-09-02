package domain

import (
	"errors"
	"strings"
	"testing"
)

func TestNewExample(t *testing.T) {
	t.Run("success with valid input", func(t *testing.T) {
		name := "John Doe"
		description := "Valid example description"
		createdBy := int64(100)

		ex, err := NewExample(name, description, createdBy)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ex == nil {
			t.Fatal("expected example instance, got nil")
		}
		if ex.Name != name {
			t.Fatalf("expected Name '%s', got '%s'", name, ex.Name)
		}
		if ex.Description != description {
			t.Fatalf("expected Description '%s', got '%s'", description, ex.Description)
		}
		if ex.CreatedBy != createdBy || ex.UpdatedBy != createdBy {
			t.Fatalf("expected CreatedBy/UpdatedBy %d, got %d/%d", createdBy, ex.CreatedBy, ex.UpdatedBy)
		}
		if ex.CreatedAt.IsZero() || ex.UpdatedAt.IsZero() {
			t.Fatal("expected CreatedAt and UpdatedAt to be set")
		}
	})

	t.Run("normalizes whitespace in name", func(t *testing.T) {
		ex, err := NewExample("   Jane    Doe   ", "Desc", 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ex.Name != "Jane Doe" {
			t.Fatalf("expected Name 'Jane Doe', got '%s'", ex.Name)
		}
	})

	t.Run("trims description whitespace", func(t *testing.T) {
		ex, err := NewExample("Jane Doe", "   trimmed description   ", 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ex.Description != "trimmed description" {
			t.Fatalf("expected trimmed description, got '%s'", ex.Description)
		}
	})

	t.Run("invalid name format - single word", func(t *testing.T) {
		_, err := NewExample("SingleWord", "Desc", 1)
		if !errors.Is(err, ExampleErrInvalidName) {
			t.Fatalf("expected ExampleErrInvalidName, got %v", err)
		}
	})

	t.Run("invalid name format - three words", func(t *testing.T) {
		_, err := NewExample("John Middle Doe", "Desc", 1)
		if !errors.Is(err, ExampleErrInvalidName) {
			t.Fatalf("expected ExampleErrInvalidName, got %v", err)
		}
	})

	t.Run("invalid name format - empty", func(t *testing.T) {
		_, err := NewExample("   ", "Desc", 1)
		if !errors.Is(err, ExampleErrInvalidName) {
			t.Fatalf("expected ExampleErrInvalidName, got %v", err)
		}
	})

	t.Run("description exceeds maximum length", func(t *testing.T) {
		tooLongDesc := strings.Repeat("a", ExampleDescriptionMaxLength+1)
		_, err := NewExample("John Doe", tooLongDesc, 1)
		if !errors.Is(err, ExampleErrDescriptionTooLong) {
			t.Fatalf("expected ExampleErrDescriptionTooLong, got %v", err)
		}
	})

	t.Run("description at exactly maximum length", func(t *testing.T) {
		maxDesc := strings.Repeat("a", ExampleDescriptionMaxLength)
		ex, err := NewExample("John Doe", maxDesc, 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ex.Description != maxDesc {
			t.Fatal("expected description to match max length")
		}
	})
}

func TestUpdateExample(t *testing.T) {
	t.Run("success update", func(t *testing.T) {
		ex, err := NewExample("John Doe", "Old Description", 1)
		if err != nil {
			t.Fatalf("unexpected setup error: %v", err)
		}

		oldUpdatedAt := ex.UpdatedAt
		updatedBy := int64(1)

		err = ex.UpdateExample("Jane Smith", "New Description", updatedBy)
		if err != nil {
			t.Fatalf("unexpected update error: %v", err)
		}

		if ex.Name != "Jane Smith" {
			t.Fatalf("expected Name 'Jane Smith', got '%s'", ex.Name)
		}
		if ex.Description != "New Description" {
			t.Fatalf("expected Description 'New Description', got '%s'", ex.Description)
		}
		if ex.UpdatedBy != updatedBy {
			t.Fatalf("expected UpdatedBy %d, got %d", updatedBy, ex.UpdatedBy)
		}
		if ex.UpdatedAt.Before(oldUpdatedAt) {
			t.Fatal("expected UpdatedAt to be updated")
		}
	})

	t.Run("invalid name format on update", func(t *testing.T) {
		ex, err := NewExample("John Doe", "Old Description", 1)
		if err != nil {
			t.Fatalf("unexpected setup error: %v", err)
		}

		err = ex.UpdateExample("InvalidNameOnly", "New Description", 1)
		if !errors.Is(err, ExampleErrInvalidName) {
			t.Fatalf("expected ExampleErrInvalidName, got %v", err)
		}
	})

	t.Run("description too long on update", func(t *testing.T) {
		ex, err := NewExample("John Doe", "Old Description", 1)
		if err != nil {
			t.Fatalf("unexpected setup error: %v", err)
		}

		tooLongDesc := strings.Repeat("b", ExampleDescriptionMaxLength+1)
		err = ex.UpdateExample("Jane Smith", tooLongDesc, 1)
		if !errors.Is(err, ExampleErrDescriptionTooLong) {
			t.Fatalf("expected ExampleErrDescriptionTooLong, got %v", err)
		}
	})
}
