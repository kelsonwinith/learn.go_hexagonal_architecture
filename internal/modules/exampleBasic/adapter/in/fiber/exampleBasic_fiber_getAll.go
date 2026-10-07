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

type ExampleBasicFiberGetAll struct {
	useCase exampleBasicDomain.ExampleBasicUsecaseGetAll
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicFiberGetAll(useCase exampleBasicDomain.ExampleBasicUsecaseGetAll) *ExampleBasicFiberGetAll {
	return &ExampleBasicFiberGetAll{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetAllExampleBasics
// @Summary Get all examples
// @Description Get all examples
// @Tags examplebasic
// @Produce json
// @Success 200 {array} exampleBasicDto.ExampleBasicResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic [get]
func (h *ExampleBasicFiberGetAll) Handle(c fiber.Ctx) error {
	_, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context())
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleBasicDto.ToExampleBasicResponses(res))
}
