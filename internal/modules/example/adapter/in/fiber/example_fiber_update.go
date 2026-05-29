package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/in/fiber/dto"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
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
// @Param example body exampleDto.UpdateExampleRequest true "Update Example"
// @Success 200 {object} exampleDto.ExampleResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/example/{id} [put]
func (h *ExampleFiberUpdate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleDto.ExampleRequestParams, sharedFiber.Empty, exampleDto.UpdateExampleRequest](c)
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	domainReq := req.Body.ToDomain()
	domainReq.ID = req.URI.ID

	res, err := h.useCase.Execute(c.Context(), domainReq)
	if err != nil {
		return sharedFiber.ErrorResponse(c, err)
	}

	return sharedFiber.SuccessResponse(c, fiber.StatusOK, exampleDto.ToExampleResponse(res))
}
