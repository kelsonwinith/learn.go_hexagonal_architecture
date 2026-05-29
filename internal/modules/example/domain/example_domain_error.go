package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

var (
	ExampleErrNotFound           = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleErrorPrefixID, "001"), "example not found")
	ExampleErrInvalidName        = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleErrorPrefixID, "002"), "example name must be in format: [First name] [Last name]")
	ExampleErrDescriptionTooLong = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleErrorPrefixID, "003"), "example description must be 255 characters or fewer")
)
