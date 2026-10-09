package dto

import (
	time "time"

	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	CreatedBy   int64     `json:"created_by"`
	UpdatedBy   int64     `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExampleProductRequestParams struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type ExampleProductGetPaginatedQuery struct {
	Page     int    `query:"page" validate:"omitempty,min=1"`
	PageSize int    `query:"page_size" validate:"omitempty,min=1"`
	Search   string `query:"search" validate:"omitempty,max=255"`
}

type ExampleProductCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
	Price       int64  `json:"price" validate:"gte=0"`
}

type ExampleProductCreateMultipleRequest struct {
	ExampleProducts []ExampleProductCreateRequest `json:"examples" validate:"required,min=1,dive"`
}

type UpdateExampleProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
	Price       int64  `json:"price" validate:"gte=0"`
}

// ============================================================================
// Methods
// ============================================================================

func (e ExampleProductCreateRequest) ToDomain(createdBy int64) exampleProductDomain.ExampleProduct {
	return exampleProductDomain.ExampleProduct{
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

func (r ExampleProductCreateMultipleRequest) ToDomain(createdBy int64) []exampleProductDomain.ExampleProduct {
	examples := make([]exampleProductDomain.ExampleProduct, len(r.ExampleProducts))
	for i, e := range r.ExampleProducts {
		examples[i] = e.ToDomain(createdBy)
	}
	return examples
}

func (e UpdateExampleProductRequest) ToDomain(updatedBy int64) exampleProductDomain.ExampleProduct {
	return exampleProductDomain.ExampleProduct{
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		UpdatedBy:   updatedBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToExampleProductResponse(e *exampleProductDomain.ExampleProduct) ExampleProductResponse {
	return ExampleProductResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		CreatedBy:   e.CreatedBy,
		UpdatedBy:   e.UpdatedBy,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToExampleProductResponses(examples []*exampleProductDomain.ExampleProduct) []ExampleProductResponse {
	res := make([]ExampleProductResponse, len(examples))
	for i, e := range examples {
		res[i] = ToExampleProductResponse(e)
	}
	return res
}
