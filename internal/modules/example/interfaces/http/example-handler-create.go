package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	dm "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

var _ = dm.Example{}

type CreateExampleHandler struct {
	useCase domain.CreateExampleUseCase
}

func NewCreateExampleHandler(useCase domain.CreateExampleUseCase) *CreateExampleHandler {
	return &CreateExampleHandler{useCase: useCase}
}

// Handle CreateExample
// @Summary Create a new example
// @Description Create a new example with the input payload
// @Tags examples
// @Accept json
// @Produce json
// @Param example body createRequest true "Create Example"
// @Success 201 {object} exampleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /examples [post]
func (h *CreateExampleHandler) Handle(c *fiber.Ctx) error {
	var req createRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	res, err := h.useCase.Execute(c.Context(), req.toDomain())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(toExampleResponse(res))
}

type createRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *createRequest) toDomain() domain.Example {
	return domain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}
