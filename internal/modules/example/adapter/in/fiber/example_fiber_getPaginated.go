package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/in/fiber/dto"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Types
// ============================================================================

type ExampleFiberGetPaginated struct {
	useCase exampleDomain.ExampleUsecaseGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleFiberGetPaginated(useCase exampleDomain.ExampleUsecaseGetPaginated) *ExampleFiberGetPaginated {
	return &ExampleFiberGetPaginated{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetPaginatedExamples
// @Summary Get examples with pagination
// @Description Get a page of examples ordered by creation date
// @Tags example
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param page_size query int false "Items per page (default 10, max 100)"
// @Param search query string false "Search by name or description"
// @Success 200 {object} sharedFiber.ResponsePaginatedData[exampleDto.ExampleResponse]
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/example/paginated [get]
func (h *ExampleFiberGetPaginated) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, exampleDto.ExampleGetPaginatedQuery, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.Query.Page, req.Query.PageSize, req.Query.Search)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponsePaginated(c, exampleDto.ToExampleResponses(res.Items), res.Pagination, res.Total)
}
