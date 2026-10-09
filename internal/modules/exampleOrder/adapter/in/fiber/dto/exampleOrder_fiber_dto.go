package dto

import (
	time "time"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderResponse struct {
	ID          string                        `json:"id"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	UserID      string                        `json:"user_id"`
	Products    []ExampleOrderProductResponse `json:"products"`
	CreatedBy   string                        `json:"created_by"`
	UpdatedBy   string                        `json:"updated_by"`
	CreatedAt   time.Time                     `json:"created_at"`
	UpdatedAt   time.Time                     `json:"updated_at"`
}

type ExampleOrderProductResponse struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	ProductID string    `json:"product_id"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExampleOrderRequestParams struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type ExampleOrderGetPaginatedQuery struct {
	Page     int    `query:"page" validate:"omitempty,min=1"`
	PageSize int    `query:"page_size" validate:"omitempty,min=1"`
	Search   string `query:"search" validate:"omitempty,max=255"`
}

type ExampleOrderCreateRequest struct {
	Name        string                             `json:"name" validate:"required,max=255"`
	Description string                             `json:"description" validate:"omitempty,max=255"`
	UserID      string                             `json:"user_id" validate:"required,uuid4"`
	Products    []ExampleOrderProductCreateRequest `json:"products" validate:"required,min=1,dive"`
}

type ExampleOrderProductCreateRequest struct {
	ProductID string `json:"product_id" validate:"required,uuid4"`
	Quantity  int    `json:"quantity" validate:"required,gt=0"`
}

// ============================================================================
// Methods
// ============================================================================

func (request ExampleOrderCreateRequest) ToDomain(createdBy string) exampleOrderDomain.ExampleOrder {
	exampleOrderProducts := make([]*exampleOrderDomain.ExampleOrderProduct, len(request.Products))
	for i, productRequest := range request.Products {
		exampleOrderProducts[i] = &exampleOrderDomain.ExampleOrderProduct{
			ProductID: productRequest.ProductID,
			Quantity:  productRequest.Quantity,
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		}
	}

	return exampleOrderDomain.ExampleOrder{
		Name:        request.Name,
		Description: request.Description,
		UserID:      request.UserID,
		Products:    exampleOrderProducts,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToExampleOrderResponse(exampleOrder *exampleOrderDomain.ExampleOrder) ExampleOrderResponse {
	return ExampleOrderResponse{
		ID:          exampleOrder.ID,
		Name:        exampleOrder.Name,
		Description: exampleOrder.Description,
		UserID:      exampleOrder.UserID,
		Products:    ToExampleOrderProductResponses(exampleOrder.Products),
		CreatedBy:   exampleOrder.CreatedBy,
		UpdatedBy:   exampleOrder.UpdatedBy,
		CreatedAt:   exampleOrder.CreatedAt,
		UpdatedAt:   exampleOrder.UpdatedAt,
	}
}

func ToExampleOrderResponses(exampleOrders []*exampleOrderDomain.ExampleOrder) []ExampleOrderResponse {
	res := make([]ExampleOrderResponse, len(exampleOrders))
	for i, exampleOrder := range exampleOrders {
		res[i] = ToExampleOrderResponse(exampleOrder)
	}

	return res
}

func ToExampleOrderProductResponses(exampleOrderProducts []*exampleOrderDomain.ExampleOrderProduct) []ExampleOrderProductResponse {
	res := make([]ExampleOrderProductResponse, len(exampleOrderProducts))
	for i, exampleOrderProduct := range exampleOrderProducts {
		res[i] = ExampleOrderProductResponse{
			ID:        exampleOrderProduct.ID,
			OrderID:   exampleOrderProduct.OrderID,
			ProductID: exampleOrderProduct.ProductID,
			Name:      exampleOrderProduct.Name,
			Quantity:  exampleOrderProduct.Quantity,
			CreatedBy: exampleOrderProduct.CreatedBy,
			UpdatedBy: exampleOrderProduct.UpdatedBy,
			CreatedAt: exampleOrderProduct.CreatedAt,
			UpdatedAt: exampleOrderProduct.UpdatedAt,
		}
	}

	return res
}
