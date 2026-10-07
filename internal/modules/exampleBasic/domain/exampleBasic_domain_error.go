package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Variables
// ============================================================================

var (
	ExampleBasicErrNotFound           = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleBasicErrPrefixID, "001"), "example not found", nil)
	ExampleBasicErrInvalidName        = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleBasicErrPrefixID, "002"), "example name must be in format: [First name] [Last name]", nil)
	ExampleBasicErrDescriptionTooLong = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleBasicErrPrefixID, "003"), "example description must be 255 characters or fewer", nil)
	ExampleBasicErrForbidden          = sharedDomain.NewError(sharedDomain.Forbidden, sharedDomain.BuildErrorID(sharedDomain.ExampleBasicErrPrefixID, "004"), "example can only be modified by its creator", nil)
)
