package sharedDomain

import (
	context "context"
)

// ============================================================================
// Types
// ============================================================================

type TokenClaims struct {
	Subject string
}

type TokenService interface {
	Generate(ctx context.Context, subject string) (string, error)
	Parse(ctx context.Context, token string) (TokenClaims, error)
}
