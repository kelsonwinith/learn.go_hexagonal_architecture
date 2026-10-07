package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Variables
// ============================================================================

var (
	ExampleErrNotFound           = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleBasicErrPrefixID, "001"), "example not found", nil)
	ExampleErrInvalidName        = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleBasicErrPrefixID, "002"), "example name must be in format: [First name] [Last name]", nil)
	ExampleErrDescriptionTooLong = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleBasicErrPrefixID, "003"), "example description must be 255 characters or fewer", nil)
	ExampleErrForbidden          = sharedDomain.NewError(sharedDomain.Forbidden, sharedDomain.BuildErrorID(sharedDomain.ExampleBasicErrPrefixID, "004"), "example can only be modified by its creator", nil)
)
