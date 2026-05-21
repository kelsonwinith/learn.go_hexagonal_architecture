package fiber

import (
	errors "errors"

	fiber "github.com/gofiber/fiber/v2"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleFiberDelete struct {
	useCase exampleDomain.ExampleUsecaseDelete
}

func NewExampleFiberDelete(useCase exampleDomain.ExampleUsecaseDelete) *ExampleFiberDelete {
	return &ExampleFiberDelete{useCase: useCase}
}

// Handle DeleteExample
// @Summary Delete an example
// @Description Delete an example by ID
// @Tags example
// @Produce json
// @Param id path string true "Example ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /example/{id} [delete]
func (h *ExampleFiberDelete) Handle(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.useCase.Execute(c.Context(), id)
	if err != nil {
		if errors.Is(err, exampleDomain.ErrExampleNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Example not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
