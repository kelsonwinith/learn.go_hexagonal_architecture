package sharedDomain

import (
	fmt "fmt"
	http "net/http"
)

type ErrorType string

const (
	ErrorTypeBadRequest     ErrorType = "BAD_REQUEST"
	ErrorTypeNotFound       ErrorType = "NOT_FOUND"
	ErrorTypeInternalServer ErrorType = "INTERNAL_SERVER_ERROR"
)

type ErrorPrefix string

const (
	SystemErrorPrefixID  ErrorPrefix = "SYSM"
	ExampleErrorPrefixID ErrorPrefix = "E"
)

type ErrorStatus struct {
	HTTPCode int
	Type     ErrorType
}

var (
	BadRequest          = ErrorStatus{HTTPCode: http.StatusBadRequest, Type: ErrorTypeBadRequest}
	NotFound            = ErrorStatus{HTTPCode: http.StatusNotFound, Type: ErrorTypeNotFound}
	InternalServerError = ErrorStatus{HTTPCode: http.StatusInternalServerError, Type: ErrorTypeInternalServer}
)

type Error struct {
	HTTPCode int
	Type     ErrorType
	ID       string
	Message  string
}

func NewError(status ErrorStatus, id, message string) *Error {
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

func BuildErrorID(prefix ErrorPrefix, id string) string {
	return string(prefix) + id
}
