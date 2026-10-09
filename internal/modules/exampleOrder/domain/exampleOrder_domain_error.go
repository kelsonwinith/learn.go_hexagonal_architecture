package domain

import (
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Variables
// ============================================================================

var (
	ExampleOrderErrNotFound               = sharedDomain.NewError(sharedDomain.NotFound, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "001"), "example order not found", nil)
	ExampleOrderErrInvalidName            = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "002"), "order name is required and must be 255 characters or fewer", nil)
	ExampleOrderErrDescriptionTooLong     = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "003"), "order description must be 255 characters or fewer", nil)
	ExampleOrderErrNoProducts             = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "004"), "order must have at least one product", nil)
	ExampleOrderErrInvalidProductName     = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "005"), "product name is required and must be 255 characters or fewer", nil)
	ExampleOrderErrInvalidProductQuantity = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "006"), "product quantity must be greater than zero", nil)
	ExampleOrderErrForbidden              = sharedDomain.NewError(sharedDomain.Forbidden, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "007"), "order can only be modified by its creator", nil)
	ExampleOrderErrInvalidUserID          = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "008"), "order user id is required", nil)
	ExampleOrderErrInvalidProductID       = sharedDomain.NewError(sharedDomain.BadRequest, sharedDomain.BuildErrorID(sharedDomain.ExampleOrderErrPrefixID, "009"), "order product id is required", nil)
)
