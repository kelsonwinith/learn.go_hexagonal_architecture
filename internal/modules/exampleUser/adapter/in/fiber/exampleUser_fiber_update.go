package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleUserDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/adapter/in/fiber/dto"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserFiberUpdate struct {
	useCase exampleUserDomain.ExampleUserUsecaseUpdate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserFiberUpdate(useCase exampleUserDomain.ExampleUserUsecaseUpdate) *ExampleUserFiberUpdate {
	return &ExampleUserFiberUpdate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle UpdateExampleUser
// @Summary Update an example
// @Description Update an example by ID
// @Tags exampleuser
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param id path string true "Example ID"
// @Param example body exampleUserDto.UpdateExampleUserRequest true "Update Example"
// @Success 200 {object} exampleUserDto.ExampleUserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleuser/{id} [put]
func (h *ExampleUserFiberUpdate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleUserDto.ExampleUserRequestParams, sharedFiber.Empty, exampleUserDto.UpdateExampleUserRequest](c)
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

	return sharedFiber.ResponseSuccess(c, exampleUserDto.ToExampleUserResponse(res))
}
