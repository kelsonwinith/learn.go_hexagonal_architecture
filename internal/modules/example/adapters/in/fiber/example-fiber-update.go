package fiber

import (
	errors "errors"

	fiber "github.com/gofiber/fiber/v2"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleFiberUpdate struct {
	useCase exampleDomain.ExampleUsecaseUpdate
}

func NewExampleFiberUpdate(useCase exampleDomain.ExampleUsecaseUpdate) *ExampleFiberUpdate {
	return &ExampleFiberUpdate{useCase: useCase}
}

// Handle UpdateExample
// @Summary Update an example
// @Description Update an example by ID
// @Tags example
// @Accept json
// @Produce json
// @Param id path string true "Example ID"
// @Param example body updateRequest true "Update Example"
// @Success 200 {object} exampleResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /example/{id} [put]
func (h *ExampleFiberUpdate) Handle(c *fiber.Ctx) error {
	id := c.Params("id")
	var req updateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	domainReq := req.toDomain()
	domainReq.ID = id

	res, err := h.useCase.Execute(c.Context(), domainReq)
	if err != nil {
		if errors.Is(err, exampleDomain.ErrExampleNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Example not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(toExampleResponse(res))
}
