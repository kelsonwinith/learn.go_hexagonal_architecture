package dto

import (
	time "time"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExampleUserLoginResponse struct {
	Token string `json:"token"`
}

type ExampleUserRequestParams struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type ExampleUserRegisterRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ExampleUserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// ============================================================================
// Methods
// ============================================================================

func (request ExampleUserRegisterRequest) ToDomain(createdBy string) exampleUserDomain.ExampleUser {
	return exampleUserDomain.ExampleUser{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}
}

// ============================================================================
// Functions
// ============================================================================

func ToExampleUserResponse(exampleUser *exampleUserDomain.ExampleUser) ExampleUserResponse {
	return ExampleUserResponse{
		ID:        exampleUser.ID,
		Name:      exampleUser.Name,
		Email:     exampleUser.Email,
		CreatedBy: exampleUser.CreatedBy,
		UpdatedBy: exampleUser.UpdatedBy,
		CreatedAt: exampleUser.CreatedAt,
		UpdatedAt: exampleUser.UpdatedAt,
	}
}
