package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Variables
// ============================================================================

var (
	ExampleUserErrNotFound           = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "001"), "example not found", nil)
	ExampleUserErrInvalidName        = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "002"), "example name must be in format: [First name] [Last name]", nil)
	ExampleUserErrDescriptionTooLong = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "003"), "example description must be 255 characters or fewer", nil)
	ExampleUserErrForbidden          = sharedDomain.NewError(sharedDomain.Forbidden, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "004"), "example can only be modified by its creator", nil)
)
