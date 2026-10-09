package password

import (
	golangBcrypt "golang.org/x/crypto/bcrypt"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserPasswordComparer struct{}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserPasswordComparer() *ExampleUserPasswordComparer {
	return &ExampleUserPasswordComparer{}
}

// ============================================================================
// Methods
// ============================================================================

func (c *ExampleUserPasswordComparer) Execute(hashedPassword, password string) bool {
	return golangBcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
