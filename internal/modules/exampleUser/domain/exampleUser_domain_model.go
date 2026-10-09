package domain

import (
	strings "strings"
	time "time"
	utf8 "unicode/utf8"
)

// ============================================================================
// Constants
// ============================================================================

const ExampleUserDescriptionMaxLength = 255

// ============================================================================
// Types
// ============================================================================

type ExampleUser struct {
	ID          string
	Name        string
	Description string
	CreatedBy   int64
	UpdatedBy   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUser(name, description string, createdBy int64) (*ExampleUser, error) {
	name, description, err := validateExampleUser(name, description)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &ExampleUser{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleUser) UpdateExampleUser(name, description string, updatedBy int64) error {
	name, description, err := validateExampleUser(name, description)
	if err != nil {
		return err
	}

	e.Name = name
	e.Description = description
	e.UpdatedBy = updatedBy
	e.UpdatedAt = time.Now().UTC()

	return nil
}

func (e *ExampleUser) IsOwnedBy(userID int64) bool {
	return e.CreatedBy == userID
}

// ============================================================================
// Functions
// ============================================================================

func validateExampleUser(name, description string) (string, string, error) {
	name, err := validateName(name)
	if err != nil {
		return "", "", err
	}

	description, err = validateDescription(description)
	if err != nil {
		return "", "", err
	}

	return name, description, nil
}

func validateName(name string) (string, error) {
	fields := strings.Fields(name)

	if len(fields) != 2 {
		return "", ExampleUserErrInvalidName
	}

	return strings.Join(fields, " "), nil
}

func validateDescription(description string) (string, error) {
	description = strings.TrimSpace(description)
	if utf8.RuneCountInString(description) > ExampleUserDescriptionMaxLength {
		return "", ExampleUserErrDescriptionTooLong
	}

	return description, nil
}
