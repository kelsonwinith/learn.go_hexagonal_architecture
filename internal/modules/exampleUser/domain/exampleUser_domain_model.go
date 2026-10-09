package domain

import (
	strings "strings"
	time "time"
	utf8 "unicode/utf8"
)

// ============================================================================
// Constants
// ============================================================================

const (
	ExampleUserNameMaxLength     = 255
	ExampleUserEmailMaxLength    = 255
	ExampleUserPasswordMinLength = 8
)

// ============================================================================
// Types
// ============================================================================

type ExampleUser struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedBy int64
	UpdatedBy int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUser(name, email, password string, createdBy int64) (*ExampleUser, error) {
	name, email, password, err := validateExampleUser(name, email, password)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &ExampleUser{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// ============================================================================
// Methods
// ============================================================================

func (u *ExampleUser) IsOwnedBy(userID int64) bool {
	return u.CreatedBy == userID
}

// ============================================================================
// Functions
// ============================================================================

func validateExampleUser(name, email, password string) (string, string, string, error) {
	name, err := validateName(name)
	if err != nil {
		return "", "", "", err
	}

	email, err = validateEmail(email)
	if err != nil {
		return "", "", "", err
	}

	password, err = validatePassword(password)
	if err != nil {
		return "", "", "", err
	}

	return name, email, password, nil
}

func validateName(name string) (string, error) {
	name = strings.TrimSpace(name)

	if name == "" || utf8.RuneCountInString(name) > ExampleUserNameMaxLength {
		return "", ExampleUserErrInvalidName
	}

	return name, nil
}

func validateEmail(email string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	if email == "" || utf8.RuneCountInString(email) > ExampleUserEmailMaxLength || !strings.Contains(email, "@") {
		return "", ExampleUserErrInvalidEmail
	}

	return email, nil
}

func validatePassword(password string) (string, error) {
	if utf8.RuneCountInString(password) < ExampleUserPasswordMinLength {
		return "", ExampleUserErrInvalidPassword
	}

	return password, nil
}
