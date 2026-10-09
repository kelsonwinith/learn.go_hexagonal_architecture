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

type ExampleUserFiberRegister struct {
	usecase exampleUserDomain.ExampleUserUsecaseRegister
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberRegister(usecase exampleUserDomain.ExampleUserUsecaseRegister) *ExampleUserFiberRegister {
	return &ExampleUserFiberRegister{usecase: usecase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle RegisterExampleUser
// @Summary Register a new example user
// @Description Register a new example user from name, email and password
// @Tags Example User
// @Accept json
// @Produce json
// @Param example body exampleUserDto.ExampleUserRegisterRequest true "Register ExampleUser"
// @Success 201 {object} exampleUserDto.ExampleUserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser/register [post]
func (h *ExampleUserFiberRegister) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleUserDto.ExampleUserRegisterRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.usecase.Execute(c.Context(), req.Body.ToDomain(""))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleUserDto.ToExampleUserResponse(res))
}
