package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleBasicDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/adapter/in/fiber/dto"
	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Types
// ============================================================================

type ExampleBasicFiberGetPaginated struct {
	useCase exampleBasicDomain.ExampleBasicUsecaseGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicFiberGetPaginated(useCase exampleBasicDomain.ExampleBasicUsecaseGetPaginated) *ExampleBasicFiberGetPaginated {
	return &ExampleBasicFiberGetPaginated{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetPaginatedExampleBasics
// @Summary Get examples with pagination
// @Description Get a page of examples ordered by creation date
// @Tags examplebasic
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param page_size query int false "Items per page (default 10, max 100)"
// @Param search query string false "Search by name or description"
// @Success 200 {object} sharedFiber.ResponsePaginatedData[exampleBasicDto.ExampleBasicResponse]
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic/paginated [get]
func (h *ExampleBasicFiberGetPaginated) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, exampleBasicDto.ExampleBasicGetPaginatedQuery, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.Query.Page, req.Query.PageSize, req.Query.Search)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponsePaginated(c, exampleBasicDto.ToExampleBasicResponses(res.Items), res.Pagination, res.Total)
}
