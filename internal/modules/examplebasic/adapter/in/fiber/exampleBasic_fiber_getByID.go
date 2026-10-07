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

type ExampleBasicFiberGetByID struct {
	useCase exampleBasicDomain.ExampleBasicUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicFiberGetByID(useCase exampleBasicDomain.ExampleBasicUsecaseGetByID) *ExampleBasicFiberGetByID {
	return &ExampleBasicFiberGetByID{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetExampleBasicByID
// @Summary Get an example by ID
// @Description Get an example by ID
// @Tags examplebasic
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} exampleBasicDto.ExampleBasicResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic/{id} [get]
func (h *ExampleBasicFiberGetByID) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleBasicDto.ExampleBasicRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.URI.ID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleBasicDto.ToExampleBasicResponse(res))
}
