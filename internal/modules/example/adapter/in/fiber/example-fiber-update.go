package fiber

import (
	fiber "github.com/gofiber/fiber/v2"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
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
// @Router /api/v1/example/{id} [put]
func (h *ExampleFiberUpdate) Handle(c *fiber.Ctx) error {
	id := c.Params("id")
	var req updateRequest
	if err := c.BodyParser(&req); err != nil {
		return sharedFiber.ErrorResponse(c, sharedDomain.New(sharedDomain.BadRequest, "E005", "invalid request body"))
	}

	domainReq := req.toDomain()
	domainReq.ID = id

	res, err := h.useCase.Execute(c.Context(), domainReq)
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	return sharedFiber.SuccessResponse(c, fiber.StatusOK, toExampleResponse(res))
}
