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

type ExampleProductFiberGetByID struct {
	useCase exampleProductDomain.ExampleProductUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductFiberGetByID(useCase exampleProductDomain.ExampleProductUsecaseGetByID) *ExampleProductFiberGetByID {
	return &ExampleProductFiberGetByID{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetExampleProductByID
// @Summary Get an example by ID
// @Description Get an example by ID
// @Tags exampleproduct
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} exampleProductDto.ExampleProductResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleproduct/{id} [get]
func (h *ExampleProductFiberGetByID) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleProductDto.ExampleProductRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.URI.ID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleProductDto.ToExampleProductResponse(res))
}
