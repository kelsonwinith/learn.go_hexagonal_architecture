package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type GetExampleByIDHandler struct {
	useCase domain.GetExampleByIDUseCase
}

func NewGetExampleByIDHandler(useCase domain.GetExampleByIDUseCase) *GetExampleByIDHandler {
	return &GetExampleByIDHandler{useCase: useCase}
}

// Handle GetExampleByID
// @Summary Get an example by ID
// @Description Get an example by ID
// @Tags examples
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} exampleResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /examples/{id} [get]
func (h *GetExampleByIDHandler) Handle(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.useCase.Execute(c.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrExampleNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Example not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(toExampleResponse(res))
}
