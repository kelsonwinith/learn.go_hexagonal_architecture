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

type ExampleBasicFiberUpdate struct {
	useCase exampleBasicDomain.ExampleBasicUsecaseUpdate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleBasicFiberUpdate(useCase exampleBasicDomain.ExampleBasicUsecaseUpdate) *ExampleBasicFiberUpdate {
	return &ExampleBasicFiberUpdate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle UpdateExampleBasic
// @Summary Update an example
// @Description Update an example by ID
// @Tags examplebasic
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param id path string true "Example ID"
// @Param example body exampleBasicDto.UpdateExampleBasicRequest true "Update Example"
// @Success 200 {object} exampleBasicDto.ExampleBasicResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/examplebasic/{id} [put]
func (h *ExampleBasicFiberUpdate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleBasicDto.ExampleBasicRequestParams, sharedFiber.Empty, exampleBasicDto.UpdateExampleBasicRequest](c)
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

	return sharedFiber.ResponseSuccess(c, exampleBasicDto.ToExampleBasicResponse(res))
}
