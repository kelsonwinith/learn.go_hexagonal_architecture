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

type ExampleProductFiberCreateMultiple struct {
	useCase exampleProductDomain.ExampleProductUsecaseCreateMultiple
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductFiberCreateMultiple(useCase exampleProductDomain.ExampleProductUsecaseCreateMultiple) *ExampleProductFiberCreateMultiple {
	return &ExampleProductFiberCreateMultiple{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateMultipleExampleProducts
// @Summary Create multiple examples in a transaction
// @Description Create multiple examples atomically - if any fails, all are rolled back
// @Tags exampleproduct
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param examples body exampleProductDto.ExampleProductCreateMultipleRequest true "Create Multiple Examples"
// @Success 201 {object} []exampleProductDto.ExampleProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleproduct/batch [post]
func (h *ExampleProductFiberCreateMultiple) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleProductDto.ExampleProductCreateMultipleRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleProductDto.ToExampleProductResponses(res))
}
