package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

var (
	ExampleErrNotFound           = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleErrPrefixID, "001"), "example not found", nil)
	ExampleErrInvalidName        = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleErrPrefixID, "002"), "example name must be in format: [First name] [Last name]", nil)
	ExampleErrDescriptionTooLong = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleErrPrefixID, "003"), "example description must be 255 characters or fewer", nil)
)
