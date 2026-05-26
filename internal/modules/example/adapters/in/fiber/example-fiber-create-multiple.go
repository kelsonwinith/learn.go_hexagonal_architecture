package fiber

import (
	fiber "github.com/gofiber/fiber/v2"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleFiberCreateMultiple struct {
	useCase exampleDomain.ExampleUsecaseCreateMultiple
}

func NewExampleFiberCreateMultiple(useCase exampleDomain.ExampleUsecaseCreateMultiple) *ExampleFiberCreateMultiple {
	return &ExampleFiberCreateMultiple{useCase: useCase}
}

// Handle CreateMultipleExamples
// @Summary Create multiple examples in a transaction
// @Description Create multiple examples atomically - if any fails, all are rolled back
// @Tags example
// @Accept json
// @Produce json
// @Param examples body createMultipleRequest true "Create Multiple Examples"
// @Success 201 {object} []exampleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/example/batch [post]
func (h *ExampleFiberCreateMultiple) Handle(c *fiber.Ctx) error {
	var req createMultipleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if len(req.Examples) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "At least one example is required"})
	}

	res, err := h.useCase.Execute(c.Context(), req.toDomain())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(toExampleResponses(res))
}
