package fiber

import (
	fiber "github.com/gofiber/fiber/v2"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
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
// @Success 200 {array} exampleResponse
// @Failure 500 {object} map[string]string
// @Router /example [get]
func (h *ExampleFiberGetAll) Handle(c *fiber.Ctx) error {
	res, err := h.useCase.Execute(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(toExampleResponses(res))
}
