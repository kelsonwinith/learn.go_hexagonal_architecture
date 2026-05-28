package sharedDomain

import (
	fmt "fmt"
	nethttp "net/http"
)

type ErrorType string

const (
	ErrorTypeBadRequest     ErrorType = "BAD_REQUEST"
	ErrorTypeNotFound       ErrorType = "NOT_FOUND"
	ErrorTypeInternalServer ErrorType = "INTERNAL_SERVER_ERROR"
)

type Status struct {
	HTTPCode int
	Type     ErrorType
}

var (
	BadRequest          = Status{HTTPCode: nethttp.StatusBadRequest, Type: ErrorTypeBadRequest}
	NotFound            = Status{HTTPCode: nethttp.StatusNotFound, Type: ErrorTypeNotFound}
	InternalServerError = Status{HTTPCode: nethttp.StatusInternalServerError, Type: ErrorTypeInternalServer}
)

type Error struct {
	HTTPCode int
	Type     ErrorType
	ID       string
	Message  string
}

func New(status Status, id, message string) *Error {
	return &Error{
		HTTPCode: status.HTTPCode,
		Type:     status.Type,
		ID:       id,
		Message:  message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.ID, e.Message)
}
