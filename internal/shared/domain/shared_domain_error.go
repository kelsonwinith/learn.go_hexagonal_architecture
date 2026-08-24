package sharedDomain

import (
	fmt "fmt"
	http "net/http"
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

const (
	ErrorTypeBadRequest     ErrorType = "BAD_REQUEST"
	ErrorTypeNotFound       ErrorType = "NOT_FOUND"
	ErrorTypeInternalServer ErrorType = "INTERNAL_SERVER_ERROR"
	ErrorTypeUnauthorized   ErrorType = "UNAUTHORIZED"
	ErrorTypeForbidden      ErrorType = "FORBIDDEN"
)

const (
	SystemErrPrefixID  ErrorPrefix = "SYS"
	FiberErrPrefixID   ErrorPrefix = "FIB"
	ExampleErrPrefixID ErrorPrefix = "ERR"
)

var (
	BadRequest    = ErrorStatus{HTTPCode: http.StatusBadRequest, Type: ErrorTypeBadRequest}
	NotFound      = ErrorStatus{HTTPCode: http.StatusNotFound, Type: ErrorTypeNotFound}
	InternalError = ErrorStatus{HTTPCode: http.StatusInternalServerError, Type: ErrorTypeInternalServer}
	Unauthorized  = ErrorStatus{HTTPCode: http.StatusUnauthorized, Type: ErrorTypeUnauthorized}
	Forbidden     = ErrorStatus{HTTPCode: http.StatusForbidden, Type: ErrorTypeForbidden}
)

var (
	SystemErrInternal = NewError(InternalError, BuildErrorID(SystemErrPrefixID, "001"), "internal server error", nil)
)

var (
	FiberErrUnauthorized      = NewError(Unauthorized, BuildErrorID(FiberErrPrefixID, "001"), "unauthorized", nil)
	FiberErrInvalidURI        = NewError(BadRequest, BuildErrorID(FiberErrPrefixID, "002"), "invalid uri", nil)
	FiberErrInvalidQuery      = NewError(BadRequest, BuildErrorID(FiberErrPrefixID, "003"), "invalid query", nil)
	FiberErrInvalidBody       = NewError(BadRequest, BuildErrorID(FiberErrPrefixID, "004"), "invalid body", nil)
	FiberErrMissingAuthHeader = NewError(Unauthorized, BuildErrorID(FiberErrPrefixID, "005"), "missing authorization header", nil)
	FiberErrInvalidAuthHeader = NewError(Unauthorized, BuildErrorID(FiberErrPrefixID, "006"), "invalid authorization header", nil)
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
