package dto

import (
	time "time"

	exampleAdvancedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/domain"
)

// ============================================================================
// Types
// ============================================================================

type ParentResponse struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Children    []ChildResponse `json:"children"`
	CreatedBy   int64           `json:"created_by"`
	UpdatedBy   int64           `json:"updated_by"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type ChildResponse struct {
	ID        string    `json:"id"`
	ParentID  string    `json:"parent_id"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	CreatedBy int64     `json:"created_by"`
	UpdatedBy int64     `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ParentRequestParams struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type ParentCreateRequest struct {
	Name        string               `json:"name" validate:"required,max=255"`
	Description string               `json:"description" validate:"omitempty,max=255"`
	Children    []ChildCreateRequest `json:"children" validate:"required,min=1,dive"`
}

type ChildCreateRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
	Quantity int    `json:"quantity" validate:"required,gt=0"`
}

// ============================================================================
// Methods
// ============================================================================

func (r ParentCreateRequest) ToDomain(createdBy int64) exampleAdvancedDomain.Parent {
	children := make([]*exampleAdvancedDomain.Child, len(r.Children))
	for i, child := range r.Children {
		children[i] = &exampleAdvancedDomain.Child{
			Name:      child.Name,
			Quantity:  child.Quantity,
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		}
	}

	return exampleAdvancedDomain.Parent{
		Name:        r.Name,
		Description: r.Description,
		Children:    children,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToParentResponse(parent *exampleAdvancedDomain.Parent) ParentResponse {
	return ParentResponse{
		ID:          parent.ID,
		Name:        parent.Name,
		Description: parent.Description,
		Children:    ToChildResponses(parent.Children),
		CreatedBy:   parent.CreatedBy,
		UpdatedBy:   parent.UpdatedBy,
		CreatedAt:   parent.CreatedAt,
		UpdatedAt:   parent.UpdatedAt,
	}
}

func ToChildResponses(children []*exampleAdvancedDomain.Child) []ChildResponse {
	res := make([]ChildResponse, len(children))
	for i, child := range children {
		res[i] = ChildResponse{
			ID:        child.ID,
			ParentID:  child.ParentID,
			Name:      child.Name,
			Quantity:  child.Quantity,
			CreatedBy: child.CreatedBy,
			UpdatedBy: child.UpdatedBy,
			CreatedAt: child.CreatedAt,
			UpdatedAt: child.UpdatedAt,
		}
	}

	return res
}
