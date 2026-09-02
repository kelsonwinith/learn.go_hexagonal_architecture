package dto

import (
	"testing"
	"time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func TestExampleCreateRequestToDomain(t *testing.T) {
	req := ExampleCreateRequest{
		Name:        "John Doe",
		Description: "A description",
	}

	domain := req.ToDomain(42)

	if domain.Name != req.Name {
		t.Fatalf("Name mismatch: got %s, want %s", domain.Name, req.Name)
	}
	if domain.Description != req.Description {
		t.Fatalf("Description mismatch: got %s, want %s", domain.Description, req.Description)
	}
	if domain.CreatedBy != 42 || domain.UpdatedBy != 42 {
		t.Fatalf("CreatedBy/UpdatedBy mismatch: got %d/%d, want 42/42", domain.CreatedBy, domain.UpdatedBy)
	}
}

func TestExampleCreateMultipleRequestToDomain(t *testing.T) {
	req := ExampleCreateMultipleRequest{
		Examples: []ExampleCreateRequest{
			{Name: "User One", Description: "Desc One"},
			{Name: "User Two", Description: "Desc Two"},
		},
	}

	domains := req.ToDomain(99)

	if len(domains) != 2 {
		t.Fatalf("expected 2 domains, got %d", len(domains))
	}
	if domains[0].Name != "User One" || domains[0].CreatedBy != 99 {
		t.Fatalf("first item mismatch: %+v", domains[0])
	}
	if domains[1].Name != "User Two" || domains[1].CreatedBy != 99 {
		t.Fatalf("second item mismatch: %+v", domains[1])
	}
}

func TestUpdateExampleRequestToDomain(t *testing.T) {
	req := UpdateExampleRequest{
		Name:        "Jane Doe",
		Description: "Updated description",
	}

	domain := req.ToDomain(77)

	if domain.Name != req.Name {
		t.Fatalf("Name mismatch: got %s, want %s", domain.Name, req.Name)
	}
	if domain.Description != req.Description {
		t.Fatalf("Description mismatch: got %s, want %s", domain.Description, req.Description)
	}
	if domain.UpdatedBy != 77 {
		t.Fatalf("UpdatedBy mismatch: got %d, want 77", domain.UpdatedBy)
	}
}

func TestToExampleResponse(t *testing.T) {
	now := time.Now().UTC()
	domain := &exampleDomain.Example{
		ID:          "uuid-123",
		Name:        "John Doe",
		Description: "Sample description",
		CreatedBy:   12,
		UpdatedBy:   34,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	res := ToExampleResponse(domain)

	if res.ID != domain.ID {
		t.Fatalf("ID mismatch: got %s, want %s", res.ID, domain.ID)
	}
	if res.Name != domain.Name {
		t.Fatalf("Name mismatch: got %s, want %s", res.Name, domain.Name)
	}
	if res.Description != domain.Description {
		t.Fatalf("Description mismatch: got %s, want %s", res.Description, domain.Description)
	}
	if res.CreatedBy != domain.CreatedBy {
		t.Fatalf("CreatedBy mismatch: got %d, want %d", res.CreatedBy, domain.CreatedBy)
	}
	if res.UpdatedBy != domain.UpdatedBy {
		t.Fatalf("UpdatedBy mismatch: got %d, want %d", res.UpdatedBy, domain.UpdatedBy)
	}
	if !res.CreatedAt.Equal(domain.CreatedAt) || !res.UpdatedAt.Equal(domain.UpdatedAt) {
		t.Fatalf("timestamp mismatch: got %v/%v, want %v/%v", res.CreatedAt, res.UpdatedAt, domain.CreatedAt, domain.UpdatedAt)
	}
}

func TestToExampleResponses(t *testing.T) {
	domains := []*exampleDomain.Example{
		{ID: "1", Name: "One"},
		{ID: "2", Name: "Two"},
	}

	responses := ToExampleResponses(domains)

	if len(responses) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(responses))
	}
	if responses[0].ID != "1" || responses[1].ID != "2" {
		t.Fatalf("unexpected IDs: %s, %s", responses[0].ID, responses[1].ID)
	}
}
