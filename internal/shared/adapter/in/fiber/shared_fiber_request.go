package fiber

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v3"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

type Validator struct {
	validate *validator.Validate
}

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

func (v *Validator) Validate(out interface{}) error {
	return v.validate.Struct(out)
}

type Empty struct{}

type Request[URI any, Query any, Body any] struct {
	URI   URI
	Query Query
	Body  Body
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
	return sharedDomain.NewError(sharedDomain.BadRequest, "E005", fmt.Sprintf("invalid request %s: %s", source, err.Error()))
}
