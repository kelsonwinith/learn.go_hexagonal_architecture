package fiber

import (
	errors "errors"

	fiber "github.com/gofiber/fiber/v3"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

var InternalServerError = sharedDomain.New(sharedDomain.InternalServerError, "SYS001", "internal server error")

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type ErrorBody struct {
	Type    sharedDomain.ErrorType `json:"type"`
	ID      string                 `json:"id"`
	Message string                 `json:"message"`
}

func SuccessResponse(c fiber.Ctx, status int, data interface{}) error {
	return SuccessResponseWithMessage(c, status, data)
}

func SuccessResponseWithMessage(c fiber.Ctx, status int, data interface{}) error {
	if status == fiber.StatusNoContent {
		return c.SendStatus(status)
	}

	return c.Status(status).JSON(Response{
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(c fiber.Ctx, err error) error {
	var appErr *sharedDomain.Error
	if !errors.As(err, &appErr) {
		appErr = InternalServerError
	}

	return c.Status(appErr.HTTPCode).JSON(Response{
		Success: false,
		Error: ErrorBody{
			Type:    appErr.Type,
			ID:      appErr.ID,
			Message: appErr.Message,
		},
	})
}
