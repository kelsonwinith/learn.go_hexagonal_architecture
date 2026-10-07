package dto

import (
	time "time"

	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   int64     `json:"created_by"`
	UpdatedBy   int64     `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExampleBasicRequestParams struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type ExampleBasicGetPaginatedQuery struct {
	Page     int    `query:"page" validate:"omitempty,min=1"`
	PageSize int    `query:"page_size" validate:"omitempty,min=1"`
	Search   string `query:"search" validate:"omitempty,max=255"`
}

type ExampleBasicCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

type ExampleBasicCreateMultipleRequest struct {
	ExampleBasics []ExampleBasicCreateRequest `json:"examples" validate:"required,min=1,dive"`
}

type UpdateExampleBasicRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

// ============================================================================
// Methods
// ============================================================================

func (e ExampleBasicCreateRequest) ToDomain(createdBy int64) exampleBasicDomain.ExampleBasic {
	return exampleBasicDomain.ExampleBasic{
		Name:        e.Name,
		Description: e.Description,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

func (r ExampleBasicCreateMultipleRequest) ToDomain(createdBy int64) []exampleBasicDomain.ExampleBasic {
	examples := make([]exampleBasicDomain.ExampleBasic, len(r.ExampleBasics))
	for i, e := range r.ExampleBasics {
		examples[i] = e.ToDomain(createdBy)
	}
	return examples
}

func (e UpdateExampleBasicRequest) ToDomain(updatedBy int64) exampleBasicDomain.ExampleBasic {
	return exampleBasicDomain.ExampleBasic{
		Name:        e.Name,
		Description: e.Description,
		UpdatedBy:   updatedBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToExampleBasicResponse(e *exampleBasicDomain.ExampleBasic) ExampleBasicResponse {
	return ExampleBasicResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatedBy:   e.CreatedBy,
		UpdatedBy:   e.UpdatedBy,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToExampleBasicResponses(examples []*exampleBasicDomain.ExampleBasic) []ExampleBasicResponse {
	res := make([]ExampleBasicResponse, len(examples))
	for i, e := range examples {
		res[i] = ToExampleBasicResponse(e)
	}
	return res
}
