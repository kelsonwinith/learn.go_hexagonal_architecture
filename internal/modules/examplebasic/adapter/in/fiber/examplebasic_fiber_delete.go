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

type ExampleFiberDelete struct {
	useCase exampleBasicDomain.ExampleUsecaseDelete
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleFiberDelete(useCase exampleBasicDomain.ExampleUsecaseDelete) *ExampleFiberDelete {
	return &ExampleFiberDelete{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle DeleteExample
// @Summary Delete an example
// @Description Delete an example by ID
// @Tags examplebasic
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param id path string true "Example ID"
// @Success 204
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic/{id} [delete]
func (h *ExampleFiberDelete) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleBasicDto.ExampleRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	err = h.useCase.Execute(c.Context(), req.URI.ID, userID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseNoContent(c)
}
