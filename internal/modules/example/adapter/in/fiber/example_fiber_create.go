package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/in/fiber/dto"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

type ExampleFiberCreate struct {
	useCase exampleDomain.ExampleUsecaseCreate
}

func NewExampleFiberCreate(useCase exampleDomain.ExampleUsecaseCreate) *ExampleFiberCreate {
	return &ExampleFiberCreate{useCase: useCase}
}

// Handle CreateExample
// @Summary Create a new example
// @Description Create a new example with the input payload
// @Tags example
// @Accept json
// @Produce json
// @Param example body exampleDto.ExampleCreateRequest true "Create Example"
// @Success 201 {object} exampleDto.ExampleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/example [post]
func (h *ExampleFiberCreate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleDto.ExampleCreateRequest](c)
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain())
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	return sharedFiber.SuccessResponse(c, fiber.StatusCreated, exampleDto.ToExampleResponse(res))
}
