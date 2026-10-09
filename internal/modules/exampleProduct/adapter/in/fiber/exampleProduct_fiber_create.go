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

type ExampleProductFiberCreate struct {
	useCase exampleProductDomain.ExampleProductUsecaseCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductFiberCreate(useCase exampleProductDomain.ExampleProductUsecaseCreate) *ExampleProductFiberCreate {
	return &ExampleProductFiberCreate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateExampleProduct
// @Summary Create a new example
// @Description Create a new example with the input payload
// @Tags exampleproduct
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param example body exampleProductDto.ExampleProductCreateRequest true "Create Example"
// @Success 201 {object} exampleProductDto.ExampleProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleproduct [post]
func (h *ExampleProductFiberCreate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleProductDto.ExampleProductCreateRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleProductDto.ToExampleProductResponse(res))
}
