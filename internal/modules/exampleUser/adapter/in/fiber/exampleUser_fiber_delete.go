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

type ExampleUserFiberDelete struct {
	useCase exampleUserDomain.ExampleUserUsecaseDelete
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberDelete(useCase exampleUserDomain.ExampleUserUsecaseDelete) *ExampleUserFiberDelete {
	return &ExampleUserFiberDelete{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle DeleteExampleUser
// @Summary Delete an example
// @Description Delete an example by ID
// @Tags exampleuser
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param id path string true "Example ID"
// @Success 204
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser/{id} [delete]
func (h *ExampleUserFiberDelete) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleUserDto.ExampleUserRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
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
