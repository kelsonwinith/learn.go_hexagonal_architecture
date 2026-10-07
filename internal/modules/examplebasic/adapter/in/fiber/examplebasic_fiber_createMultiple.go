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

type ExampleFiberCreateMultiple struct {
	useCase exampleBasicDomain.ExampleUsecaseCreateMultiple
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleFiberCreateMultiple(useCase exampleBasicDomain.ExampleUsecaseCreateMultiple) *ExampleFiberCreateMultiple {
	return &ExampleFiberCreateMultiple{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateMultipleExamples
// @Summary Create multiple examples in a transaction
// @Description Create multiple examples atomically - if any fails, all are rolled back
// @Tags examplebasic
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param examples body exampleBasicDto.ExampleCreateMultipleRequest true "Create Multiple Examples"
// @Success 201 {object} []exampleBasicDto.ExampleResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic/batch [post]
func (h *ExampleFiberCreateMultiple) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleBasicDto.ExampleCreateMultipleRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleBasicDto.ToExampleResponses(res))
}
