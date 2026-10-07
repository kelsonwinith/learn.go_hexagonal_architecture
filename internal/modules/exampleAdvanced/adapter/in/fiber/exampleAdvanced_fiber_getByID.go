package fiber

import (
	fiber "github.com/gofiber/fiber/v3"
	exampleAdvancedDto "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/adapter/in/fiber/dto"
	exampleAdvancedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/domain"
	sharedFiber "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/adapter/in/fiber"
)

// ============================================================================
// Types
// ============================================================================

type ExampleAdvancedFiberGetByID struct {
	useCase exampleAdvancedDomain.ExampleAdvancedUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleAdvancedFiberGetByID(useCase exampleAdvancedDomain.ExampleAdvancedUsecaseGetByID) *ExampleAdvancedFiberGetByID {
	return &ExampleAdvancedFiberGetByID{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle GetExampleAdvancedParentByID
// @Summary Get an example advanced parent by ID
// @Description Returns a parent aggregate with its children preloaded
// @Tags exampleadvanced
// @Produce json
// @Param id path string true "Parent ID"
// @Success 200 {object} exampleAdvancedDto.ParentResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleadvanced/{id} [get]
func (h *ExampleAdvancedFiberGetByID) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[exampleAdvancedDto.ParentRequestParams, sharedFiber.Empty, sharedFiber.Empty](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	res, err := h.useCase.Execute(c.Context(), req.URI.ID)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseSuccess(c, exampleAdvancedDto.ToParentResponse(res))
}
