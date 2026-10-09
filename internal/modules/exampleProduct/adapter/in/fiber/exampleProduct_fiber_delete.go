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

type ExampleProductFiberDelete struct {
	usecase exampleProductDomain.ExampleProductUsecaseDelete
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProductFiberDelete(usecase exampleProductDomain.ExampleProductUsecaseDelete) *ExampleProductFiberDelete {
	return &ExampleProductFiberDelete{usecase: usecase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle DeleteExampleProduct
// @Summary Delete an example
// @Description Delete an example by ID
// @Tags Example Product
// @Produce json
// @Security BearerAuth
// @Param id path string true "Example ID"
// @Success 204
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleproduct/{id} [delete]
func (h *ExampleProductFiberDelete) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleProductDto.ExampleProductRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	err = h.usecase.Execute(c.Context(), req.URI.ID, userID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseNoContent(c)
}
