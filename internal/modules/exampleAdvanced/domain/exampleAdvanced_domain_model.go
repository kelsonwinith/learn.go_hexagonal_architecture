package domain

import (
	strings "strings"
	time "time"
	utf8 "unicode/utf8"
)

// ============================================================================
// Constants
// ============================================================================

const (
	ParentNameMaxLength        = 255
	ParentDescriptionMaxLength = 255
	ChildNameMaxLength         = 255
	ChildQuantityMin           = 1
)

// ============================================================================
// Types
// ============================================================================

type Parent struct {
	ID          string
	Name        string
	Description string
	Children    []*Child
	CreatedBy   int64
	UpdatedBy   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Child struct {
	ID        string
	ParentID  string
	Name      string
	Quantity  int
	CreatedBy int64
	UpdatedBy int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ============================================================================
// Constructors
// ============================================================================

func NewParent(name, description string, children []*Child, createdBy int64) (*Parent, error) {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > ParentNameMaxLength {
		return nil, ExampleAdvancedErrInvalidName
	}

	description = strings.TrimSpace(description)
	if utf8.RuneCountInString(description) > ParentDescriptionMaxLength {
		return nil, ExampleAdvancedErrDescriptionTooLong
	}

	if len(children) == 0 {
		return nil, ExampleAdvancedErrNoChildren
	}

	validatedChildren := make([]*Child, len(children))
	for i, child := range children {
		validated, err := NewChild(child.Name, child.Quantity, createdBy)
		if err != nil {
			return nil, err
		}
		validatedChildren[i] = validated
	}

	now := time.Now().UTC()

	return &Parent{
		Name:        name,
		Description: description,
		Children:    validatedChildren,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func NewChild(name string, quantity int, createdBy int64) (*Child, error) {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > ChildNameMaxLength {
		return nil, ExampleAdvancedErrInvalidChildName
	}

	if quantity < ChildQuantityMin {
		return nil, ExampleAdvancedErrInvalidChildQuantity
	}

	now := time.Now().UTC()

	return &Child{
		Name:      name,
		Quantity:  quantity,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// ============================================================================
// Methods
// ============================================================================

func (p *Parent) IsOwnedBy(userID int64) bool {
	return p.CreatedBy == userID
}
