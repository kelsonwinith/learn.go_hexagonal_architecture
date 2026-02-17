package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/application"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type UpdateExampleHandler struct {
	useCase *application.UpdateExampleUseCase
}

func NewUpdateExampleHandler(useCase *application.UpdateExampleUseCase) *UpdateExampleHandler {
	return &UpdateExampleHandler{useCase: useCase}
}

// Handle UpdateExample
// @Summary Update an example
// @Description Update an example by ID
// @Tags examples
// @Accept json
// @Produce json
// @Param id path string true "Example ID"
// @Param example body application.UpdateExampleInput true "Update Example"
// @Success 200 {object} domain.Example
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /examples/{id} [put]
func (h *UpdateExampleHandler) Handle(c *fiber.Ctx) error {
	id := c.Params("id")
	var req application.UpdateExampleInput
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	req.ID = id

	res, err := h.useCase.Execute(c.Context(), req)
	if err != nil {
		if errors.Is(err, domain.ErrExampleNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Example not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
