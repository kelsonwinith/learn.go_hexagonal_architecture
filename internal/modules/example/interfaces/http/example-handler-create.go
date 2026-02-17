package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/application"
	dm "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

var _ = dm.Example{}

type CreateExampleHandler struct {
	useCase *application.CreateExampleUseCase
}

func NewCreateExampleHandler(useCase *application.CreateExampleUseCase) *CreateExampleHandler {
	return &CreateExampleHandler{useCase: useCase}
}

// Handle CreateExample
// @Summary Create a new example
// @Description Create a new example with the input payload
// @Tags examples
// @Accept json
// @Produce json
// @Param example body application.CreateExampleInput true "Create Example"
// @Success 201 {object} dm.Example
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /examples [post]
func (h *CreateExampleHandler) Handle(c *fiber.Ctx) error {
	var req application.CreateExampleInput
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	res, err := h.useCase.Execute(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
