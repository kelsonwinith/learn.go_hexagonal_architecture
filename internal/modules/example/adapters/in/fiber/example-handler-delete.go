package fiber

import (
	errors "errors"

	fiber "github.com/gofiber/fiber/v2"
	domain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type DeleteExampleHandler struct {
	useCase domain.DeleteExampleUseCase
}

func NewDeleteExampleHandler(useCase domain.DeleteExampleUseCase) *DeleteExampleHandler {
	return &DeleteExampleHandler{useCase: useCase}
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
func (h *DeleteExampleHandler) Handle(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.useCase.Execute(c.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrExampleNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Example not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
