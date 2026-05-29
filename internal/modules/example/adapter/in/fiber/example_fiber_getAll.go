package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/in/fiber/dto"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

type ExampleFiberGetAll struct {
	useCase exampleDomain.ExampleUsecaseGetAll
}

func NewExampleFiberGetAll(useCase exampleDomain.ExampleUsecaseGetAll) *ExampleFiberGetAll {
	return &ExampleFiberGetAll{useCase: useCase}
}

// Handle GetAllExamples
// @Summary Get all examples
// @Description Get all examples
// @Tags example
// @Produce json
// @Success 200 {array} exampleDto.ExampleResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/example [get]
func (h *ExampleFiberGetAll) Handle(c fiber.Ctx) error {
	_, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	res, err := h.useCase.Execute(c.Context())
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	return sharedFiber.SuccessResponse(c, fiber.StatusOK, exampleDto.ToExampleResponses(res))
}
