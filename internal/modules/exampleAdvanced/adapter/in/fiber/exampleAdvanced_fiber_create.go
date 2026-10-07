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

type ExampleAdvancedFiberCreate struct {
	useCase exampleAdvancedDomain.ExampleAdvancedUsecaseCreate
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleAdvancedFiberCreate(useCase exampleAdvancedDomain.ExampleAdvancedUsecaseCreate) *ExampleAdvancedFiberCreate {
	return &ExampleAdvancedFiberCreate{useCase: useCase}
}

// ============================================================================
// Methods
// ============================================================================

// Handle CreateExampleAdvancedParent
// @Summary Create an example advanced parent with children
// @Description Demonstrates nested aggregate creation, a transaction, and a second output port
// @Tags exampleadvanced
// @Accept json
// @Produce json
// @Security UserIdAuth
// @Param example-user-id header int true "Authenticated User ID"
// @Param parent body exampleAdvancedDto.ParentCreateRequest true "Create Parent With Children"
// @Success 201 {object} exampleAdvancedDto.ParentResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/exampleadvanced [post]
func (h *ExampleAdvancedFiberCreate) Handle(c fiber.Ctx) error {
	req, err := sharedFiber.Bind[sharedFiber.Empty, sharedFiber.Empty, exampleAdvancedDto.ParentCreateRequest](c)
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	userID := sharedFiber.GetAuthUserID(c)
	res, err := h.useCase.Execute(c.Context(), req.Body.ToDomain(userID))
	if err != nil {
		return sharedFiber.ResponseError(c, err)
	}

	return sharedFiber.ResponseCreated(c, exampleAdvancedDto.ToParentResponse(res))
}
