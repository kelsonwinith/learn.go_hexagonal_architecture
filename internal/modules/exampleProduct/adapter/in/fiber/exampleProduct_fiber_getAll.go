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

type ExampleProductFiberGetAll struct {
	useCase exampleProductDomain.ExampleProductUsecaseGetAll
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductFiberGetAll(useCase exampleProductDomain.ExampleProductUsecaseGetAll) *ExampleProductFiberGetAll {
	return &ExampleProductFiberGetAll{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetAllExampleProducts
// @Summary Get all examples
// @Description Get all examples
// @Tags exampleproduct
// @Produce json
// @Success 200 {array} exampleProductDto.ExampleProductResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleproduct [get]
func (h *ExampleProductFiberGetAll) Handle(c fiber.Ctx) error {
	_, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context())
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleProductDto.ToExampleProductResponses(res))
}
