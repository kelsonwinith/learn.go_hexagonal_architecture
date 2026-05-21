package fiber

import (
	errors "errors"

	fiber "github.com/gofiber/fiber/v2"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleFiberGetByID struct {
	useCase exampleDomain.ExampleUsecaseGetByID
}

func NewExampleFiberGetByID(useCase exampleDomain.ExampleUsecaseGetByID) *ExampleFiberGetByID {
	return &ExampleFiberGetByID{useCase: useCase}
}

// Handle GetExampleByID
// @Summary Get an example by ID
// @Description Get an example by ID
// @Tags example
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} exampleResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /example/{id} [get]
func (h *ExampleFiberGetByID) Handle(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.useCase.Execute(c.Context(), id)
	if err != nil {
		if errors.Is(err, exampleDomain.ErrExampleNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Example not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(toExampleResponse(res))
}
