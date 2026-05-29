package dto

import (
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToExampleResponse(e *exampleDomain.Example) ExampleResponse {
	return ExampleResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
	}
}

func ToExampleResponses(examples []*exampleDomain.Example) []ExampleResponse {
	res := make([]ExampleResponse, len(examples))
	for i, e := range examples {
		res[i] = ToExampleResponse(e)
	}
	return res
}

type CreateExampleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e CreateExampleRequest) ToDomain() exampleDomain.Example {
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}

type CreateMultipleExampleRequest struct {
	Examples []CreateExampleRequest `json:"examples"`
}

func (r CreateMultipleExampleRequest) ToDomain() []exampleDomain.Example {
	examples := make([]exampleDomain.Example, len(r.Examples))
	for i, e := range r.Examples {
		examples[i] = e.ToDomain()
	}
	return examples
}

type UpdateExampleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *UpdateExampleRequest) ToDomain() exampleDomain.Example {
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}
