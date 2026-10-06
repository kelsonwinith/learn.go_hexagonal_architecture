package dto

import (
	time "time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   int64     `json:"created_by"`
	UpdatedBy   int64     `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExampleRequestParams struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type ExampleGetPaginatedQuery struct {
	Page     int    `query:"page" validate:"omitempty,min=1"`
	PageSize int    `query:"page_size" validate:"omitempty,min=1"`
	Search   string `query:"search" validate:"omitempty,max=255"`
}

type ExampleCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

type ExampleCreateMultipleRequest struct {
	Examples []ExampleCreateRequest `json:"examples" validate:"required,min=1,dive"`
}

type UpdateExampleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

// ============================================================================
// Methods
// ============================================================================

func (e ExampleCreateRequest) ToDomain(createdBy int64) exampleDomain.Example {
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

func (r ExampleCreateMultipleRequest) ToDomain(createdBy int64) []exampleDomain.Example {
	examples := make([]exampleDomain.Example, len(r.Examples))
	for i, e := range r.Examples {
		examples[i] = e.ToDomain(createdBy)
	}
	return examples
}

func (e UpdateExampleRequest) ToDomain(updatedBy int64) exampleDomain.Example {
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
		UpdatedBy:   updatedBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToExampleResponse(e *exampleDomain.Example) ExampleResponse {
	return ExampleResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatedBy:   e.CreatedBy,
		UpdatedBy:   e.UpdatedBy,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToExampleResponses(examples []*exampleDomain.Example) []ExampleResponse {
	res := make([]ExampleResponse, len(examples))
	for i, e := range examples {
		res[i] = ToExampleResponse(e)
	}
	return res
}
