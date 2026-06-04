package dto

import (
	time "time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExampleRequestParams struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type ExampleCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

type ExampleCreateMultipleRequest struct {
	Examples []ExampleCreateRequest `json:"examples" validate:"required,min=1,dive"`
}

type UpdateExampleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

func ToExampleResponse(e *exampleDomain.Example) ExampleResponse {
	return ExampleResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToExampleResponses(examples []*exampleDomain.Example) []ExampleResponse {
	res := make([]ExampleResponse, len(examples))
	for i, e := range examples {
		res[i] = ToExampleResponse(e)
	}
	return res
}

func (e ExampleCreateRequest) ToDomain() exampleDomain.Example {
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}

func (r ExampleCreateMultipleRequest) ToDomain() []exampleDomain.Example {
	examples := make([]exampleDomain.Example, len(r.Examples))
	for i, e := range r.Examples {
		examples[i] = e.ToDomain()
	}
	return examples
}

func (e *UpdateExampleRequest) ToDomain() exampleDomain.Example {
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}
