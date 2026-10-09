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

type ExampleUserFiberLogin struct {
	usecase exampleUserDomain.ExampleUserUsecaseLogin
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberLogin(usecase exampleUserDomain.ExampleUserUsecaseLogin) *ExampleUserFiberLogin {
	return &ExampleUserFiberLogin{usecase: usecase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle LoginExampleUser
// @Summary Login an example user
// @Description Mock login: verifies email and password and returns a mock token
// @Tags Example User
// @Accept json
// @Produce json
// @Param example body exampleUserDto.ExampleUserLoginRequest true "Login ExampleUser"
// @Success 200 {object} exampleUserDto.ExampleUserLoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser/login [post]
func (h *ExampleUserFiberLogin) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleUserDto.ExampleUserLoginRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	token, err := h.usecase.Execute(c.Context(), req.Body.Email, req.Body.Password)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleUserDto.ExampleUserLoginResponse{Token: token})
}
