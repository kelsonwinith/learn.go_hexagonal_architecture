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

type ExampleOrderFiberGetByID struct {
	useCase exampleOrderDomain.ExampleOrderUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderFiberGetByID(useCase exampleOrderDomain.ExampleOrderUsecaseGetByID) *ExampleOrderFiberGetByID {
	return &ExampleOrderFiberGetByID{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetExampleOrderByID
// @Summary Get an example order by ID
// @Description Returns a order aggregate with its products preloaded
// @Tags Example Order
// @Produce json
// @Param id path string true "ExampleOrder ID"
// @Success 200 {object} exampleOrderDto.ExampleOrderResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleorder/{id} [get]
func (h *ExampleOrderFiberGetByID) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleOrderDto.ExampleOrderRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.URI.ID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleOrderDto.ToExampleOrderResponse(res))
}
