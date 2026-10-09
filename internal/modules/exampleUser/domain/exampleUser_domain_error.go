package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Variables
// ============================================================================

var (
	ExampleUserErrNotFound           = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "001"), "example user not found", nil)
	ExampleUserErrInvalidName        = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "002"), "user name is required and must be 255 characters or fewer", nil)
	ExampleUserErrInvalidEmail       = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "003"), "user email is required and must be a valid email address", nil)
	ExampleUserErrInvalidPassword    = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "004"), "user password must be at least 8 characters", nil)
	ExampleUserErrInvalidCredentials = sharedDomain.NewError(sharedDomain.Unauthorized, sharedDomain.BuildErrorID(sharedDomain.ExampleUserErrPrefixID, "005"), "invalid email or password", nil)
)
