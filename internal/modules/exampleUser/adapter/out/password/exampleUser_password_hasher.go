package password

import (
	golangBcrypt "golang.org/x/crypto/bcrypt"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserPasswordHasher struct{}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPasswordHasher() *ExampleUserPasswordHasher {
	return &ExampleUserPasswordHasher{}
}

// ============================================================================
// Methods
// ============================================================================

func (h *ExampleUserPasswordHasher) Execute(password string) (string, error) {
	hashed, err := golangBcrypt.GenerateFromPassword([]byte(password), golangBcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
