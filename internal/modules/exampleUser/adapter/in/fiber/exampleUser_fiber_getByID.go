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

type ExampleUserFiberGetByID struct {
	usecase exampleUserDomain.ExampleUserUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberGetByID(usecase exampleUserDomain.ExampleUserUsecaseGetByID) *ExampleUserFiberGetByID {
	return &ExampleUserFiberGetByID{usecase: usecase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetExampleUserByID
// @Summary Get an example user by ID
// @Description Get an example user by ID
// @Tags Example User
// @Produce json
// @Param id path string true "Example User ID"
// @Success 200 {object} exampleUserDto.ExampleUserResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser/{id} [get]
func (h *ExampleUserFiberGetByID) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleUserDto.ExampleUserRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.usecase.Execute(c.Context(), req.URI.ID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleUserDto.ToExampleUserResponse(res))
}
