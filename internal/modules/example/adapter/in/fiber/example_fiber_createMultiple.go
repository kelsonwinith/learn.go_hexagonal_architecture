package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/in/fiber/dto"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

type ExampleFiberCreateMultiple struct {
	useCase exampleDomain.ExampleUsecaseCreateMultiple
}

func NewExampleFiberCreateMultiple(useCase exampleDomain.ExampleUsecaseCreateMultiple) *ExampleFiberCreateMultiple {
	return &ExampleFiberCreateMultiple{useCase: useCase}
}

// Handle CreateMultipleExamples
// @Summary Create multiple examples in a transaction
// @Description Create multiple examples atomically - if any fails, all are rolled back
// @Tags example
// @Accept json
// @Produce json
// @Param examples body exampleDto.ExampleCreateMultipleRequest true "Create Multiple Examples"
// @Success 201 {object} []exampleDto.ExampleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/example/batch [post]
func (h *ExampleFiberCreateMultiple) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleDto.ExampleCreateMultipleRequest](c)
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain())
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	return sharedFiber.SuccessResponse(c, fiber.StatusCreated, exampleDto.ToExampleResponses(res))
}
