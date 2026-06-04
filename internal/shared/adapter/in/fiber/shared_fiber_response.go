package fiber

import (
	"errors"

	fiber "github.com/gofiber/fiber/v3"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

type responseBase struct {
	Success bool                   `json:"success"`
	Data    any                    `json:"data"`
	Error   *responseBaseErrorBody `json:"error"`
}

type responseBaseErrorBody struct {
	Type    sharedDomain.ErrorType `json:"type"`
	ID      string                 `json:"id"`
	Message string                 `json:"message"`
	Detail  any                    `json:"detail,omitempty"`
}

// 2XX
func ResponseSuccess(c fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(responseBase{
		Success: true,
		Data:    data,
	})
}

func ResponseCreated(c fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(responseBase{
		Success: true,
		Data:    data,
	})
}

func ResponseNoContent(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// 4XX-5XX
func ResponseError(c fiber.Ctx, err error) error {
	var appErr *sharedDomain.Error
	if !errors.As(err, &appErr) {
		appErr = sharedDomain.SystemErrInternal
	}

	return c.Status(appErr.HTTPCode).JSON(responseBase{
		Success: false,
		Error: &responseBaseErrorBody{
			Type:    appErr.Type,
			ID:      appErr.ID,
			Message: appErr.Message,
			Detail:  appErr.Detail,
		},
	})
}
