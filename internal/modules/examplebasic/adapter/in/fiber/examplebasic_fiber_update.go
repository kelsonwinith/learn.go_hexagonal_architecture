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

type ExampleFiberUpdate struct {
	useCase exampleBasicDomain.ExampleUsecaseUpdate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleFiberUpdate(useCase exampleBasicDomain.ExampleUsecaseUpdate) *ExampleFiberUpdate {
	return &ExampleFiberUpdate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle UpdateExample
// @Summary Update an example
// @Description Update an example by ID
// @Tags examplebasic
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param id path string true "Example ID"
// @Param example body exampleBasicDto.UpdateExampleRequest true "Update Example"
// @Success 200 {object} exampleBasicDto.ExampleResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic/{id} [put]
func (h *ExampleFiberUpdate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleBasicDto.ExampleRequestParams, sharedFiber.Empty, exampleBasicDto.UpdateExampleRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	domainReq := req.Body.ToDomain(userID)
	domainReq.ID = req.URI.ID

	res, err := h.useCase.Execute(c.Context(), domainReq)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleBasicDto.ToExampleResponse(res))
}
