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

type ExampleBasicFiberCreate struct {
	useCase exampleBasicDomain.ExampleBasicUsecaseCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicFiberCreate(useCase exampleBasicDomain.ExampleBasicUsecaseCreate) *ExampleBasicFiberCreate {
	return &ExampleBasicFiberCreate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateExampleBasic
// @Summary Create a new example
// @Description Create a new example with the input payload
// @Tags examplebasic
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param example body exampleBasicDto.ExampleBasicCreateRequest true "Create Example"
// @Success 201 {object} exampleBasicDto.ExampleBasicResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic [post]
func (h *ExampleBasicFiberCreate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleBasicDto.ExampleBasicCreateRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleBasicDto.ToExampleBasicResponse(res))
}
