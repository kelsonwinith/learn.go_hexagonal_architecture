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

type ExampleFiberCreate struct {
	useCase exampleBasicDomain.ExampleUsecaseCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleFiberCreate(useCase exampleBasicDomain.ExampleUsecaseCreate) *ExampleFiberCreate {
	return &ExampleFiberCreate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateExample
// @Summary Create a new example
// @Description Create a new example with the input payload
// @Tags examplebasic
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param example body exampleBasicDto.ExampleCreateRequest true "Create Example"
// @Success 201 {object} exampleBasicDto.ExampleResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic [post]
func (h *ExampleFiberCreate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleBasicDto.ExampleCreateRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleBasicDto.ToExampleResponse(res))
}
