package domain

import (
	strings "strings"
	time "time"
	utf8 "unicode/utf8"
)

const ExampleDescriptionMaxLength = 255

type Example struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewExample(name, description string) (*Example, error) {
	name, description, err := validateExample(name, description)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &Example{
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (e *Example) UpdateExample(name, description string) error {
	name, description, err := validateExample(name, description)
	if err != nil {
		return err
	}

	e.Name = name
	e.Description = description
	e.UpdatedAt = time.Now().UTC()

	return nil
}

func validateExample(name, description string) (string, string, error) {
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
		return "", ExampleErrInvalidName
	}

	return strings.Join(fields, " "), nil
}

func validateDescription(description string) (string, error) {
	description = strings.TrimSpace(description)
	if utf8.RuneCountInString(description) > ExampleDescriptionMaxLength {
		return "", ExampleErrDescriptionTooLong
	}

	return description, nil
}
