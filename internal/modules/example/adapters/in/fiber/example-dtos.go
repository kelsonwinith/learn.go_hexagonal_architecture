package fiber

import (
	domain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type exampleResponse struct {
	ID          string `json:"id" example:"uuid"`
	Name        string `json:"name" example:"Example Name"`
	Description string `json:"description" example:"Example Description"`
}

func toExampleResponse(e *domain.Example) exampleResponse {
	return exampleResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
	}
}

func toExampleResponses(examples []*domain.Example) []exampleResponse {
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

func (e createRequest) toDomain() domain.Example {
	return domain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}

type createMultipleRequest struct {
	Examples []createRequest `json:"examples"`
}

func (r createMultipleRequest) toDomain() []domain.Example {
	examples := make([]domain.Example, len(r.Examples))
	for i, e := range r.Examples {
		examples[i] = e.toDomain()
	}
	return examples
}

type updateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *updateRequest) toDomain() domain.Example {
	return domain.Example{
		Name:        e.Name,
		Description: e.Description,
	}
}
