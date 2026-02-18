package http

import "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"

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
	dtos := make([]exampleResponse, len(examples))
	for i, e := range examples {
		dtos[i] = toExampleResponse(e)
	}
	return dtos
}
