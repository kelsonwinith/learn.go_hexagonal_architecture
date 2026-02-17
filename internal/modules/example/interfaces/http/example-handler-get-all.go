package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/application"
	dm "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

var _ = dm.Example{}

type GetAllExamplesHandler struct {
	useCase *application.GetAllExamplesUseCase
}

func NewGetAllExamplesHandler(useCase *application.GetAllExamplesUseCase) *GetAllExamplesHandler {
	return &GetAllExamplesHandler{useCase: useCase}
}

// Handle GetAllExamples
// @Summary Get all examples
// @Description Get all examples
// @Tags examples
// @Produce json
// @Success 200 {array} dm.Example
// @Failure 500 {object} map[string]string
// @Router /examples [get]
func (h *GetAllExamplesHandler) Handle(c *fiber.Ctx) error {
	res, err := h.useCase.Execute(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
