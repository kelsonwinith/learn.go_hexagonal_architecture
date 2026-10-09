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

type ExampleUserFiberGetAll struct {
	useCase exampleUserDomain.ExampleUserUsecaseGetAll
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberGetAll(useCase exampleUserDomain.ExampleUserUsecaseGetAll) *ExampleUserFiberGetAll {
	return &ExampleUserFiberGetAll{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetAllExampleUsers
// @Summary Get all examples
// @Description Get all examples
// @Tags exampleuser
// @Produce json
// @Success 200 {array} exampleUserDto.ExampleUserResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser [get]
func (h *ExampleUserFiberGetAll) Handle(c fiber.Ctx) error {
	_, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context())
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleUserDto.ToExampleUserResponses(res))
}
