package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/ports"
)

type GetAllExamplesHandler struct {
	useCase ports.GetAllExamplesUseCase
}

func NewGetAllExamplesHandler(useCase ports.GetAllExamplesUseCase) *GetAllExamplesHandler {
	return &GetAllExamplesHandler{useCase: useCase}
}

// Handle GetAllExamples
// @Summary Get all examples
// @Description Get all examples
// @Tags examples
// @Produce json
// @Success 200 {array} exampleResponse
// @Failure 500 {object} map[string]string
// @Router /examples [get]
func (h *GetAllExamplesHandler) Handle(c *fiber.Ctx) error {
	res, err := h.useCase.Execute(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(toExampleResponses(res))
}
