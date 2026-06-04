package fiber

import (
	"errors"
	reflect "reflect"
	strings "strings"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v3"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

type Validator struct {
	validate *validator.Validate
}

type Request[URI any, Query any, Body any] struct {
	URI   URI
	Query Query
	Body  Body
}

type Empty struct{}

func NewValidator() *Validator {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		for _, tag := range []string{"json", "query", "params", "uri"} {
			name := strings.SplitN(field.Tag.Get(tag), ",", 2)[0]
			if name != "" && name != "-" {
				return name
			}
		}

		return field.Name
	})

	return &Validator{validate: validate}
}

func (v *Validator) Validate(out any) error {
	return v.validate.Struct(out)
}

func Bind[URI any, Query any, Body any](c fiber.Ctx) (*Request[URI, Query, Body], error) {
	var req Request[URI, Query, Body]

	if !isEmpty[URI]() {
		if err := c.Bind().URI(&req.URI); err != nil {
			return nil, invalidRequest("uri", err)
		}
	}

	if !isEmpty[Query]() {
		if err := c.Bind().Query(&req.Query); err != nil {
			return nil, invalidRequest("query", err)
		}
	}

	if !isEmpty[Body]() {
		if err := c.Bind().Body(&req.Body); err != nil {
			return nil, invalidRequest("body", err)
		}
	}

	return &req, nil
}

func isEmpty[T any]() bool {
	return reflect.TypeOf((*T)(nil)).Elem() == reflect.TypeOf(Empty{})
}

func invalidRequest(source string, err error) error {
	detail := validationErrorDetail(err)
	if len(detail) == 0 {
		detail = []string{err.Error()}
	}

	appErr := invalidRequestError(source)
	return &sharedDomain.Error{
		HTTPCode: appErr.HTTPCode,
		Type:     appErr.Type,
		ID:       appErr.ID,
		Message:  appErr.Message,
		Detail:   detail,
	}
}

func invalidRequestError(source string) *sharedDomain.Error {
	switch source {
	case "uri":
		return sharedDomain.FiberErrInvalidURI
	case "query":
		return sharedDomain.FiberErrInvalidQuery
	case "body":
		return sharedDomain.FiberErrInvalidBody
	default:
		return sharedDomain.SystemErrInternal
	}
}

func validationErrorDetail(err error) []string {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return nil
	}

	detail := make([]string, len(validationErrors))
	for i, fieldErr := range validationErrors {
		detail[i] = fieldErr.Error()
	}

	return detail
}
