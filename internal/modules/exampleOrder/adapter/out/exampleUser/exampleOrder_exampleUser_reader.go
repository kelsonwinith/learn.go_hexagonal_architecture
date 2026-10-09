package exampleUser

import (
	context "context"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderUserReader struct {
	userGetByID exampleUserDomain.ExampleUserUsecaseGetByID
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderUserReader(userGetByID exampleUserDomain.ExampleUserUsecaseGetByID) *ExampleOrderUserReader {
	return &ExampleOrderUserReader{userGetByID: userGetByID}
}

// ============================================================================
// Methods
// ============================================================================

func (r *ExampleOrderUserReader) Execute(ctx context.Context, id string) (exampleOrderDomain.ExampleOrderUser, error) {
	user, err := r.userGetByID.Execute(ctx, id)
	if err != nil {
		return exampleOrderDomain.ExampleOrderUser{}, err
	}

	return exampleOrderDomain.ExampleOrderUser{ID: user.ID, Name: user.Name}, nil
}
