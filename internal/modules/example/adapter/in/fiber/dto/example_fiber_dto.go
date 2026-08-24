package dto

import (
	time "time"

	exampleDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/example/domain"
)

type ExampleResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   int64     `json:"created_by"`
	UpdatedBy   int64     `json:"updated_by"`
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
		CreatedBy:   e.CreatedBy,
		UpdatedBy:   e.UpdatedBy,
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

func (e ExampleCreateRequest) ToDomain(createdBy ...int64) exampleDomain.Example {
	var uid int64
	if len(createdBy) > 0 {
		uid = createdBy[0]
	}
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
		CreatedBy:   uid,
		UpdatedBy:   uid,
	}
}

func (r ExampleCreateMultipleRequest) ToDomain(createdBy ...int64) []exampleDomain.Example {
	var uid int64
	if len(createdBy) > 0 {
		uid = createdBy[0]
	}
	examples := make([]exampleDomain.Example, len(r.Examples))
	for i, e := range r.Examples {
		examples[i] = e.ToDomain(uid)
	}
	return examples
}

func (e *UpdateExampleRequest) ToDomain(updatedBy ...int64) exampleDomain.Example {
	var uid int64
	if len(updatedBy) > 0 {
		uid = updatedBy[0]
	}
	return exampleDomain.Example{
		Name:        e.Name,
		Description: e.Description,
		UpdatedBy:   uid,
	}
}
