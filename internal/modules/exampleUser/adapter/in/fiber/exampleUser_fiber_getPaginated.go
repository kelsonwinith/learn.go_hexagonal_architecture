package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleUserDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/in/fiber/dto"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserFiberGetPaginated struct {
	useCase exampleUserDomain.ExampleUserUsecaseGetPaginated
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberGetPaginated(useCase exampleUserDomain.ExampleUserUsecaseGetPaginated) *ExampleUserFiberGetPaginated {
	return &ExampleUserFiberGetPaginated{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetPaginatedExampleUsers
// @Summary Get examples with pagination
// @Description Get a page of examples ordered by creation date
// @Tags exampleuser
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param page_size query int false "Items per page (default 10, max 100)"
// @Param search query string false "Search by name or description"
// @Success 200 {object} sharedFiber.ResponsePaginatedData[exampleUserDto.ExampleUserResponse]
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser/paginated [get]
func (h *ExampleUserFiberGetPaginated) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, exampleUserDto.ExampleUserGetPaginatedQuery, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.Query.Page, req.Query.PageSize, req.Query.Search)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponsePaginated(c, exampleUserDto.ToExampleUserResponses(res.Items), res.Pagination, res.Total)
}
