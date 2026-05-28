package fiber

import (
	fiber "github.com/gofiber/fiber/v2"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
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
// @Param example body createRequest true "Create Example"
// @Success 201 {object} exampleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/example [post]
func (h *ExampleFiberCreate) Handle(c *fiber.Ctx) error {
	var req createRequest
	if err := c.BodyParser(&req); err != nil {
		return sharedFiber.ErrorResponse(c, sharedDomain.New(sharedDomain.BadRequest, "E005", "invalid request body"))
	}

	res, err := h.useCase.Execute(c.Context(), req.toDomain())
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	return sharedFiber.SuccessResponse(c, fiber.StatusCreated, toExampleResponse(res))
}
