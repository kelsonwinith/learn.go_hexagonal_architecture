package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Variables
// ============================================================================

var (
	ExampleAdvancedErrNotFound             = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleAdvancedErrPrefixID, "001"), "example advanced parent not found", nil)
	ExampleAdvancedErrInvalidName          = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleAdvancedErrPrefixID, "002"), "parent name is required and must be 255 characters or fewer", nil)
	ExampleAdvancedErrDescriptionTooLong   = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleAdvancedErrPrefixID, "003"), "parent description must be 255 characters or fewer", nil)
	ExampleAdvancedErrNoChildren           = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleAdvancedErrPrefixID, "004"), "parent must have at least one child", nil)
	ExampleAdvancedErrInvalidChildName     = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleAdvancedErrPrefixID, "005"), "child name is required and must be 255 characters or fewer", nil)
	ExampleAdvancedErrInvalidChildQuantity = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleAdvancedErrPrefixID, "006"), "child quantity must be greater than zero", nil)
	ExampleAdvancedErrForbidden            = sharedDomain.NewError(sharedDomain.Forbidden, sharedDomain.BuildErrorID(sharedDomain.ExampleAdvancedErrPrefixID, "007"), "parent can only be modified by its creator", nil)
)
