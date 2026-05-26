package fiber

import (
	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type exampleResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func toExampleResponse(e *exampleDomain.Example) exampleResponse {
	return exampleResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
	}
}

func toExampleResponses(examples []*exampleDomain.Example) []exampleResponse {
	res := make([]exampleResponse, len(examples))
	for i, e := range examples {
		res[i] = toExampleResponse(e)
	}
	return res
}

type createRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e createRequest) toDomain() exampleDomain.Example {
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}

type createMultipleRequest struct {
	Examples []createRequest `json:"examples"`
}

func (r createMultipleRequest) toDomain() []exampleDomain.Example {
	examples := make([]exampleDomain.Example, len(r.Examples))
	for i, e := range r.Examples {
		examples[i] = e.toDomain()
	}
	return examples
}

type updateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *updateRequest) toDomain() exampleDomain.Example {
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}
