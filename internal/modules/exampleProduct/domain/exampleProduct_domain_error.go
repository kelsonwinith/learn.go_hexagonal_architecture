package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Variables
// ============================================================================

var (
	ExampleProductErrNotFound           = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleProductErrPrefixID, "001"), "example product not found", nil)
	ExampleProductErrInvalidName        = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleProductErrPrefixID, "002"), "product name is required and must be 255 characters or fewer", nil)
	ExampleProductErrDescriptionTooLong = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleProductErrPrefixID, "003"), "product description must be 255 characters or fewer", nil)
	ExampleProductErrInvalidPrice       = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleProductErrPrefixID, "004"), "product price must be zero or greater", nil)
	ExampleProductErrForbidden          = sharedDomain.NewError(sharedDomain.Forbidden, sharedDomain.BuildErrorID(sharedDomain.ExampleProductErrPrefixID, "005"), "product can only be modified by its creator", nil)
)
