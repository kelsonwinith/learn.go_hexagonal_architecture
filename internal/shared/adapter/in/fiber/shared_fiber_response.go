package fiber

import (
	errors "errors"

	fiber "github.com/gofiber/fiber/v3"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

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

type responsePaginatedData[T any] struct {
	Items    []T   `json:"items"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

// ============================================================================
// Functions
// ============================================================================

// 2XX
func ResponseSuccess(c fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(responseBase{
		Success: true,
		Data:    data,
	})
}

func ResponsePaginated[T any](c fiber.Ctx, items []T, pagination sharedDomain.Pagination, total int64) error {
	return c.Status(fiber.StatusOK).JSON(responseBase{
		Success: true,
		Data: responsePaginatedData[T]{
			Items:    items,
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
			Total:    total,
		},
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
