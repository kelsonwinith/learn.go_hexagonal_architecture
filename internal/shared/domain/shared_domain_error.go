package sharedDomain

import (
	fmt "fmt"
	http "net/http"
)

const (
	ErrorTypeBadRequest     ErrorType = "BAD_REQUEST"
	ErrorTypeNotFound       ErrorType = "NOT_FOUND"
	ErrorTypeInternalServer ErrorType = "INTERNAL_SERVER_ERROR"
)

const (
	SystemErrPrefixID  ErrorPrefix = "SYSM"
	FiberErrPrefixID   ErrorPrefix = "FIB"
	ExampleErrPrefixID ErrorPrefix = "E"
)

type ErrorType string

type ErrorPrefix string

type ErrorStatus struct {
	HTTPCode int
	Type     ErrorType
}

type Error struct {
	HTTPCode int
	Type     ErrorType
	ID       string
	Message  string
	Detail   any
}

var (
	BadRequest    = ErrorStatus{HTTPCode: http.StatusBadRequest, Type: ErrorTypeBadRequest}
	NotFound      = ErrorStatus{HTTPCode: http.StatusNotFound, Type: ErrorTypeNotFound}
	InternalError = ErrorStatus{HTTPCode: http.StatusInternalServerError, Type: ErrorTypeInternalServer}
)

var (
	SystemErrInternal = NewError(InternalError, BuildErrorID(SystemErrPrefixID, "001"), "internal server error", nil)
)

var (
	FiberErrInvalidURI   = NewError(BadRequest, BuildErrorID(FiberErrPrefixID, "002"), "invalid uri", nil)
	FiberErrInvalidQuery = NewError(BadRequest, BuildErrorID(FiberErrPrefixID, "003"), "invalid query", nil)
	FiberErrInvalidBody  = NewError(BadRequest, BuildErrorID(FiberErrPrefixID, "004"), "invalid body", nil)
)

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.ID, e.Message)
}

func NewError(status ErrorStatus, id, message string, detail any) *Error {
	return &Error{
		HTTPCode: status.HTTPCode,
		Type:     status.Type,
		ID:       id,
		Message:  message,
		Detail:   detail,
	}
}

func BuildErrorID(prefix ErrorPrefix, id string) string {
	return string(prefix) + id
}
