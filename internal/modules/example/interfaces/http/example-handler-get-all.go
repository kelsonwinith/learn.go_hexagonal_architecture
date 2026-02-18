package http

import (
	"github.com/gofiber/fiber/v2"
	dm "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

var _ = dm.Example{}

type GetAllExamplesHandler struct {
	useCase dm.GetAllExamplesUseCase
}

func NewGetAllExamplesHandler(useCase dm.GetAllExamplesUseCase) *GetAllExamplesHandler {
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
