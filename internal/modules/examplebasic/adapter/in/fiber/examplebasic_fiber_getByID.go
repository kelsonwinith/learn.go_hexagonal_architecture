package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleBasicDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/adapter/in/fiber/dto"
	exampleBasicDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleBasic/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Types
// ============================================================================

type ExampleFiberGetByID struct {
	useCase exampleBasicDomain.ExampleUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleFiberGetByID(useCase exampleBasicDomain.ExampleUsecaseGetByID) *ExampleFiberGetByID {
	return &ExampleFiberGetByID{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetExampleByID
// @Summary Get an example by ID
// @Description Get an example by ID
// @Tags examplebasic
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} exampleBasicDto.ExampleResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic/{id} [get]
func (h *ExampleFiberGetByID) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleBasicDto.ExampleRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.URI.ID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleBasicDto.ToExampleResponse(res))
}
