package mapper

import (
	"testing"
	"time"

	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

func TestToExampleModel(t *testing.T) {
	now := time.Now().UTC()
	domain := &exampleDomain.Example{
		ID:          "test-uuid-123",
		Name:        "John Doe",
		Description: "Test Description",
		CreatedBy:   42,
		UpdatedBy:   42,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	model := ToExampleModel(domain)

	if model.ID != domain.ID {
		t.Fatalf("ID mismatch: got %s, want %s", model.ID, domain.ID)
	}
	if model.Name != domain.Name {
		t.Fatalf("Name mismatch: got %s, want %s", model.Name, domain.Name)
	}
	if model.Description != domain.Description {
		t.Fatalf("Description mismatch: got %s, want %s", model.Description, domain.Description)
	}
	if model.CreatedBy != domain.CreatedBy {
		t.Fatalf("CreatedBy mismatch: got %d, want %d", model.CreatedBy, domain.CreatedBy)
	}
	if model.UpdatedBy != domain.UpdatedBy {
		t.Fatalf("UpdatedBy mismatch: got %d, want %d", model.UpdatedBy, domain.UpdatedBy)
	}
	if !model.CreatedAt.Equal(domain.CreatedAt) {
		t.Fatalf("CreatedAt mismatch: got %v, want %v", model.CreatedAt, domain.CreatedAt)
	}
	if !model.UpdatedAt.Equal(domain.UpdatedAt) {
		t.Fatalf("UpdatedAt mismatch: got %v, want %v", model.UpdatedAt, domain.UpdatedAt)
	}
}

func TestToExampleModels(t *testing.T) {
	now := time.Now().UTC()
	domains := []*exampleDomain.Example{
		{ID: "1", Name: "User One", CreatedBy: 10, UpdatedBy: 10, CreatedAt: now, UpdatedAt: now},
		{ID: "2", Name: "User Two", CreatedBy: 20, UpdatedBy: 20, CreatedAt: now, UpdatedAt: now},
	}

	models := ToExampleModels(domains)

	if len(models) != 2 {
		t.Fatalf("expected 2 models, got %d", len(models))
	}
	if models[0].ID != "1" || models[1].ID != "2" {
		t.Fatalf("unexpected model IDs: %s, %s", models[0].ID, models[1].ID)
	}
}

func TestToExampleDomain(t *testing.T) {
	now := time.Now().UTC()
	model := &postgresqlModel.ExampleModel{
		BaseModel: defaultModel.BaseModel{
			ID:        "db-uuid-456",
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: 99,
			UpdatedBy: 99,
		},
		Name:        "Jane Smith",
		Description: "DB Description",
	}

	domain := ToExampleDomain(model)

	if domain.ID != model.ID {
		t.Fatalf("ID mismatch: got %s, want %s", domain.ID, model.ID)
	}
	if domain.Name != model.Name {
		t.Fatalf("Name mismatch: got %s, want %s", domain.Name, model.Name)
	}
	if domain.Description != model.Description {
		t.Fatalf("Description mismatch: got %s, want %s", domain.Description, model.Description)
	}
	if domain.CreatedBy != model.CreatedBy {
		t.Fatalf("CreatedBy mismatch: got %d, want %d", domain.CreatedBy, model.CreatedBy)
	}
	if domain.UpdatedBy != model.UpdatedBy {
		t.Fatalf("UpdatedBy mismatch: got %d, want %d", domain.UpdatedBy, model.UpdatedBy)
	}
	if !domain.CreatedAt.Equal(model.CreatedAt) {
		t.Fatalf("CreatedAt mismatch: got %v, want %v", domain.CreatedAt, model.CreatedAt)
	}
	if !domain.UpdatedAt.Equal(model.UpdatedAt) {
		t.Fatalf("UpdatedAt mismatch: got %v, want %v", domain.UpdatedAt, model.UpdatedAt)
	}
}

func TestToExampleDomains(t *testing.T) {
	now := time.Now().UTC()
	models := []*postgresqlModel.ExampleModel{
		{BaseModel: defaultModel.BaseModel{ID: "1", CreatedAt: now, UpdatedAt: now}, Name: "One"},
		{BaseModel: defaultModel.BaseModel{ID: "2", CreatedAt: now, UpdatedAt: now}, Name: "Two"},
	}

	domains := ToExampleDomains(models)

	if len(domains) != 2 {
		t.Fatalf("expected 2 domains, got %d", len(domains))
	}
	if domains[0].ID != "1" || domains[1].ID != "2" {
		t.Fatalf("unexpected domain IDs: %s, %s", domains[0].ID, domains[1].ID)
	}
}
