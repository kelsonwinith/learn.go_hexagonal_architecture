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

type ExampleUserFiberCreateMultiple struct {
	useCase exampleUserDomain.ExampleUserUsecaseCreateMultiple
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberCreateMultiple(useCase exampleUserDomain.ExampleUserUsecaseCreateMultiple) *ExampleUserFiberCreateMultiple {
	return &ExampleUserFiberCreateMultiple{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateMultipleExampleUsers
// @Summary Create multiple examples in a transaction
// @Description Create multiple examples atomically - if any fails, all are rolled back
// @Tags exampleuser
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param examples body exampleUserDto.ExampleUserCreateMultipleRequest true "Create Multiple Examples"
// @Success 201 {object} []exampleUserDto.ExampleUserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser/batch [post]
func (h *ExampleUserFiberCreateMultiple) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleUserDto.ExampleUserCreateMultipleRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleUserDto.ToExampleUserResponses(res))
}
