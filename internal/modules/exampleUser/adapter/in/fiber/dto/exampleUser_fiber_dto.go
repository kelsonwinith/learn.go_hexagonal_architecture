package dto

import (
	time "time"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   int64     `json:"created_by"`
	UpdatedBy   int64     `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExampleUserRequestParams struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type ExampleUserGetPaginatedQuery struct {
	Page     int    `query:"page" validate:"omitempty,min=1"`
	PageSize int    `query:"page_size" validate:"omitempty,min=1"`
	Search   string `query:"search" validate:"omitempty,max=255"`
}

type ExampleUserCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

type ExampleUserCreateMultipleRequest struct {
	ExampleUsers []ExampleUserCreateRequest `json:"examples" validate:"required,min=1,dive"`
}

type UpdateExampleUserRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

// ============================================================================
// Methods
// ============================================================================

func (e ExampleUserCreateRequest) ToDomain(createdBy int64) exampleUserDomain.ExampleUser {
	return exampleUserDomain.ExampleUser{
		Name:        e.Name,
		Description: e.Description,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

func (r ExampleUserCreateMultipleRequest) ToDomain(createdBy int64) []exampleUserDomain.ExampleUser {
	examples := make([]exampleUserDomain.ExampleUser, len(r.ExampleUsers))
	for i, e := range r.ExampleUsers {
		examples[i] = e.ToDomain(createdBy)
	}
	return examples
}

func (e UpdateExampleUserRequest) ToDomain(updatedBy int64) exampleUserDomain.ExampleUser {
	return exampleUserDomain.ExampleUser{
		Name:        e.Name,
		Description: e.Description,
		UpdatedBy:   updatedBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToExampleUserResponse(e *exampleUserDomain.ExampleUser) ExampleUserResponse {
	return ExampleUserResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatedBy:   e.CreatedBy,
		UpdatedBy:   e.UpdatedBy,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToExampleUserResponses(examples []*exampleUserDomain.ExampleUser) []ExampleUserResponse {
	res := make([]ExampleUserResponse, len(examples))
	for i, e := range examples {
		res[i] = ToExampleUserResponse(e)
	}
	return res
}
