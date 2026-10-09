package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleProductDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/adapter/in/fiber/dto"
	exampleProductDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleProduct/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Types
// ============================================================================

type ExampleProductFiberGetPaginated struct {
	usecase exampleProductDomain.ExampleProductUsecaseGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductFiberGetPaginated(usecase exampleProductDomain.ExampleProductUsecaseGetPaginated) *ExampleProductFiberGetPaginated {
	return &ExampleProductFiberGetPaginated{usecase: usecase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetPaginatedExampleProducts
// @Summary Get example products with pagination
// @Description Get a page of example products ordered by creation date
// @Tags Example Product
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param page_size query int false "Items per page (default 10, max 100)"
// @Param search query string false "Search by name or description"
// @Success 200 {object} sharedFiber.ResponsePaginatedData[exampleProductDto.ExampleProductResponse]
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleproduct/paginated [get]
func (h *ExampleProductFiberGetPaginated) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, exampleProductDto.ExampleProductGetPaginatedQuery, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.usecase.Execute(c.Context(), req.Query.Page, req.Query.PageSize, req.Query.Search)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponsePaginated(c, exampleProductDto.ToExampleProductResponses(res.Items), res.Pagination, res.Total)
}
