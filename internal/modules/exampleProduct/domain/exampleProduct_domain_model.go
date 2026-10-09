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
	ExampleProductNameMaxLength        = 255
	ExampleProductDescriptionMaxLength = 255
	ExampleProductMinPrice             = 0
)

// ============================================================================
// Types
// ============================================================================

type ExampleProduct struct {
	ID          string
	Name        string
	Description string
	Price       int64
	CreatedBy   int64
	UpdatedBy   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleProduct(name, description string, price int64, createdBy int64) (*ExampleProduct, error) {
	name, description, price, err := validateExampleProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &ExampleProduct{
		Name:        name,
		Description: description,
		Price:       price,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// ============================================================================
// Methods
// ============================================================================

func (e *ExampleProduct) UpdateExampleProduct(name, description string, price int64, updatedBy int64) error {
	name, description, price, err := validateExampleProduct(name, description, price)
	if err != nil {
		return err
	}

	e.Name = name
	e.Description = description
	e.Price = price
	e.UpdatedBy = updatedBy
	e.UpdatedAt = time.Now().UTC()

	return nil
}

func (e *ExampleProduct) IsOwnedBy(userID int64) bool {
	return e.CreatedBy == userID
}

// ============================================================================
// Functions
// ============================================================================

func validateExampleProduct(name, description string, price int64) (string, string, int64, error) {
	name, err := validateName(name)
	if err != nil {
		return "", "", 0, err
	}

	description, err = validateDescription(description)
	if err != nil {
		return "", "", 0, err
	}

	price, err = validatePrice(price)
	if err != nil {
		return "", "", 0, err
	}

	return name, description, price, nil
}

func validateName(name string) (string, error) {
	name = strings.TrimSpace(name)

	if name == "" || utf8.RuneCountInString(name) > ExampleProductNameMaxLength {
		return "", ExampleProductErrInvalidName
	}

	return name, nil
}

func validateDescription(description string) (string, error) {
	description = strings.TrimSpace(description)
	if utf8.RuneCountInString(description) > ExampleProductDescriptionMaxLength {
		return "", ExampleProductErrDescriptionTooLong
	}

	return description, nil
}

func validatePrice(price int64) (int64, error) {
	if price < ExampleProductMinPrice {
		return 0, ExampleProductErrInvalidPrice
	}

	return price, nil
}
