package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/adapter/in/fiber/dto"
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

type ExampleFiberGetByID struct {
	useCase exampleDomain.ExampleUsecaseGetByID
}

func NewExampleFiberGetByID(useCase exampleDomain.ExampleUsecaseGetByID) *ExampleFiberGetByID {
	return &ExampleFiberGetByID{useCase: useCase}
}

// Handle GetExampleByID
// @Summary Get an example by ID
// @Description Get an example by ID
// @Tags example
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} exampleDto.ExampleResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/example/{id} [get]
func (h *ExampleFiberGetByID) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleDto.ExampleRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.URI.ID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleDto.ToExampleResponse(res))
}
