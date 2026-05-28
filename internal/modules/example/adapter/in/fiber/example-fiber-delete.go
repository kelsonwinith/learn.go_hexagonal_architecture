package fiber

import (
	fiber "github.com/gofiber/fiber/v2"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
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
// @Router /api/v1/example/{id} [delete]
func (h *ExampleFiberDelete) Handle(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.useCase.Execute(c.Context(), id)
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	return sharedFiber.SuccessResponse(c, fiber.StatusNoContent, nil)
}
