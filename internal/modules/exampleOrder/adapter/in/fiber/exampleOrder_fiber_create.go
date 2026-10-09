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

type ExampleOrderFiberCreate struct {
	useCase exampleOrderDomain.ExampleOrderUsecaseCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderFiberCreate(useCase exampleOrderDomain.ExampleOrderUsecaseCreate) *ExampleOrderFiberCreate {
	return &ExampleOrderFiberCreate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateExampleOrder
// @Summary Create an example order with products
// @Description Demonstrates nested aggregate creation, a transaction, and a second output port
// @Tags Example Order
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param order body exampleOrderDto.ExampleOrderCreateRequest true "Create ExampleOrder With Products"
// @Success 201 {object} exampleOrderDto.ExampleOrderResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleorder [post]
func (h *ExampleOrderFiberCreate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleOrderDto.ExampleOrderCreateRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleOrderDto.ToExampleOrderResponse(res))
}
