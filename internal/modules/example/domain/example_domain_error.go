package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

var (
	ExampleErrNotFound           = sharedDomain.New(sharedDomain.NotFound, "E001", "example not found")
	ExampleErrInvalidName        = sharedDomain.New(sharedDomain.BadRequest, "E002", "example name must be in format: [First name] [Last name]")
	ExampleErrDescriptionTooLong = sharedDomain.New(sharedDomain.BadRequest, "E003", "example description must be 255 characters or fewer")
)
