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
	Description string                             `json:"description" validate:"omproductpty,max=255"`
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

func (r ExampleOrderCreateRequest) ToDomain(createdBy string) exampleOrderDomain.ExampleOrder {
	products := make([]*exampleOrderDomain.ExampleOrderProduct, len(r.Products))
	for i, orderProduct := range r.Products {
		products[i] = &exampleOrderDomain.ExampleOrderProduct{
			ProductID: orderProduct.ProductID,
			Quantity:  orderProduct.Quantity,
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		}
	}

	return exampleOrderDomain.ExampleOrder{
		Name:        r.Name,
		Description: r.Description,
		UserID:      r.UserID,
		Products:    products,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToExampleOrderResponse(order *exampleOrderDomain.ExampleOrder) ExampleOrderResponse {
	return ExampleOrderResponse{
		ID:          order.ID,
		Name:        order.Name,
		Description: order.Description,
		UserID:      order.UserID,
		Products:    ToExampleOrderProductResponses(order.Products),
		CreatedBy:   order.CreatedBy,
		UpdatedBy:   order.UpdatedBy,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}

func ToExampleOrderResponses(orders []*exampleOrderDomain.ExampleOrder) []ExampleOrderResponse {
	res := make([]ExampleOrderResponse, len(orders))
	for i, order := range orders {
		res[i] = ToExampleOrderResponse(order)
	}

	return res
}

func ToExampleOrderProductResponses(products []*exampleOrderDomain.ExampleOrderProduct) []ExampleOrderProductResponse {
	res := make([]ExampleOrderProductResponse, len(products))
	for i, orderProduct := range products {
		res[i] = ExampleOrderProductResponse{
			ID:        orderProduct.ID,
			OrderID:   orderProduct.OrderID,
			ProductID: orderProduct.ProductID,
			Name:      orderProduct.Name,
			Quantity:  orderProduct.Quantity,
			CreatedBy: orderProduct.CreatedBy,
			UpdatedBy: orderProduct.UpdatedBy,
			CreatedAt: orderProduct.CreatedAt,
			UpdatedAt: orderProduct.UpdatedAt,
		}
	}

	return res
}
