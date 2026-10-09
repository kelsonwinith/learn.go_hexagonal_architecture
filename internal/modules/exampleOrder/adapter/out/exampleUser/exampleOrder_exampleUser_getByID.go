package exampleUser

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserModuleGetByID struct {
	userGetByID exampleUserDomain.ExampleUserUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserModuleGetByID(userGetByID exampleUserDomain.ExampleUserUsecaseGetByID) *ExampleUserModuleGetByID {
	return &ExampleUserModuleGetByID{userGetByID: userGetByID}
}

// ============================================================================
// Methods
// ============================================================================

func (r *ExampleUserModuleGetByID) Execute(ctx context.Context, id string) (exampleOrderDomain.ExampleOrderUser, error) {
	user, err := r.userGetByID.Execute(ctx, id)
	if err != nil {
		return exampleOrderDomain.ExampleOrderUser{}, err
	}

	return exampleOrderDomain.ExampleOrderUser{ID: user.ID, Name: user.Name}, nil
}
