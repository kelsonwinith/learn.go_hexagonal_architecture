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
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
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

type UpdateExampleProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
	Price       int64  `json:"price" validate:"gte=0"`
}

// ============================================================================
// Methods
// ============================================================================

func (request ExampleProductCreateRequest) ToDomain(createdBy string) exampleProductDomain.ExampleProduct {
	return exampleProductDomain.ExampleProduct{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

func (request UpdateExampleProductRequest) ToDomain(updatedBy string) exampleProductDomain.ExampleProduct {
	return exampleProductDomain.ExampleProduct{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		UpdatedBy:   updatedBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToExampleProductResponse(exampleProduct *exampleProductDomain.ExampleProduct) ExampleProductResponse {
	return ExampleProductResponse{
		ID:          exampleProduct.ID,
		Name:        exampleProduct.Name,
		Description: exampleProduct.Description,
		Price:       exampleProduct.Price,
		CreatedBy:   exampleProduct.CreatedBy,
		UpdatedBy:   exampleProduct.UpdatedBy,
		CreatedAt:   exampleProduct.CreatedAt,
		UpdatedAt:   exampleProduct.UpdatedAt,
	}
}

func ToExampleProductResponses(exampleProducts []*exampleProductDomain.ExampleProduct) []ExampleProductResponse {
	res := make([]ExampleProductResponse, len(exampleProducts))
	for i, exampleProduct := range exampleProducts {
		res[i] = ToExampleProductResponse(exampleProduct)
	}
	return res
}
