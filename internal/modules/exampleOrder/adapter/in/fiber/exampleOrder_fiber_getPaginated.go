package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleOrderDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/adapter/in/fiber/dto"
	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderFiberGetPaginated struct {
	usecase exampleOrderDomain.ExampleOrderUsecaseGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderFiberGetPaginated(usecase exampleOrderDomain.ExampleOrderUsecaseGetPaginated) *ExampleOrderFiberGetPaginated {
	return &ExampleOrderFiberGetPaginated{usecase: usecase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetPaginatedExampleOrders
// @Summary Get example orders with pagination
// @Description Get a page of example orders ordered by creation date
// @Tags Example Order
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param page_size query int false "Items per page (default 10, max 100)"
// @Param search query string false "Search by name or description"
// @Success 200 {object} sharedFiber.ResponsePaginatedData[exampleOrderDto.ExampleOrderResponse]
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleorder/paginated [get]
func (h *ExampleOrderFiberGetPaginated) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, exampleOrderDto.ExampleOrderGetPaginatedQuery, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.usecase.Execute(c.Context(), req.Query.Page, req.Query.PageSize, req.Query.Search)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponsePaginated(c, exampleOrderDto.ToExampleOrderResponses(res.Items), res.Pagination, res.Total)
}
