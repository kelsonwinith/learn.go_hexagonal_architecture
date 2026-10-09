package jwt

import (
	context "context"
	time "time"

	golangJwt "github.com/golang-jwt/jwt/v5"
	sharedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/shared/domain"
)

// ============================================================================
// Types
// ============================================================================

type JwtService struct {
	secret []byte
	expiry time.Duration
}

// ============================================================================
// Constructors
// ============================================================================

func NewJwtService(secret string, expiryHours int) *JwtService {
	return &JwtService{
		secret: []byte(secret),
		expiry: time.Duration(expiryHours) * time.Hour,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (j *JwtService) Generate(ctx context.Context, subject string) (string, error) {
	now := time.Now().UTC()

	claims := golangJwt.RegisteredClaims{
		Subject:   subject,
		IssuedAt:  golangJwt.NewNumericDate(now),
		ExpiresAt: golangJwt.NewNumericDate(now.Add(j.expiry)),
	}

	token := golangJwt.NewWithClaims(golangJwt.SigningMethodHS256, claims)

	return token.SignedString(j.secret)
}

func (j *JwtService) Parse(ctx context.Context, token string) (sharedDomain.TokenClaims, error) {
	claims := &golangJwt.RegisteredClaims{}

	parsed, err := golangJwt.ParseWithClaims(token, claims, func(t *golangJwt.Token) (any, error) {
		return j.secret, nil
	}, golangJwt.WithValidMethods([]string{golangJwt.SigningMethodHS256.Alg()}))
	if err != nil || !parsed.Valid {
		return sharedDomain.TokenClaims{}, sharedDomain.FiberErrUnauthorized
	}

	return sharedDomain.TokenClaims{Subject: claims.Subject}, nil
}
