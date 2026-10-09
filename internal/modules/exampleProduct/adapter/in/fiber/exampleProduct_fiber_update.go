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

type ExampleProductFiberUpdate struct {
	useCase exampleProductDomain.ExampleProductUsecaseUpdate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductFiberUpdate(useCase exampleProductDomain.ExampleProductUsecaseUpdate) *ExampleProductFiberUpdate {
	return &ExampleProductFiberUpdate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle UpdateExampleProduct
// @Summary Update an example
// @Description Update an example by ID
// @Tags Example Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Example ID"
// @Param example body exampleProductDto.UpdateExampleProductRequest true "Update Example"
// @Success 200 {object} exampleProductDto.ExampleProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleproduct/{id} [put]
func (h *ExampleProductFiberUpdate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleProductDto.ExampleProductRequestParams, sharedFiber.Empty, exampleProductDto.UpdateExampleProductRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	domainReq := req.Body.ToDomain(userID)
	domainReq.ID = req.URI.ID

	res, err := h.useCase.Execute(c.Context(), domainReq)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleProductDto.ToExampleProductResponse(res))
}
