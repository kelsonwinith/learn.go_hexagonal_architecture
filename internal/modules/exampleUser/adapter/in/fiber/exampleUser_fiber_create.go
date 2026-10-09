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

type ExampleUserFiberCreate struct {
	useCase exampleUserDomain.ExampleUserUsecaseCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberCreate(useCase exampleUserDomain.ExampleUserUsecaseCreate) *ExampleUserFiberCreate {
	return &ExampleUserFiberCreate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateExampleUser
// @Summary Create a new example
// @Description Create a new example with the input payload
// @Tags exampleuser
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param example body exampleUserDto.ExampleUserCreateRequest true "Create Example"
// @Success 201 {object} exampleUserDto.ExampleUserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser [post]
func (h *ExampleUserFiberCreate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleUserDto.ExampleUserCreateRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleUserDto.ToExampleUserResponse(res))
}
