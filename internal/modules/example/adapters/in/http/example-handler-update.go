package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	"github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/ports"
)

type UpdateExampleHandler struct {
	useCase ports.UpdateExampleUseCase
}

func NewUpdateExampleHandler(useCase ports.UpdateExampleUseCase) *UpdateExampleHandler {
	return &UpdateExampleHandler{useCase: useCase}
}

// Handle UpdateExample
// @Summary Update an example
// @Description Update an example by ID
// @Tags examples
// @Accept json
// @Produce json
// @Param id path string true "Example ID"
// @Param example body updateRequest true "Update Example"
// @Success 200 {object} exampleResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /examples/{id} [put]
func (h *UpdateExampleHandler) Handle(c *fiber.Ctx) error {
	id := c.Params("id")
	var req updateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	domainReq := req.toDomain()
	domainReq.ID = id

	res, err := h.useCase.Execute(c.Context(), domainReq)
	if err != nil {
		if errors.Is(err, domain.ErrExampleNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Example not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(toExampleResponse(res))
}

type updateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *updateRequest) toDomain() domain.Example {
	return domain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}
