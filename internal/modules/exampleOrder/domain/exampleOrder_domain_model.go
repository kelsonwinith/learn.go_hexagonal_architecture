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
	ExampleOrderNameMaxLength        = 255
	ExampleOrderDescriptionMaxLength = 255
	ExampleOrderProductNameMaxLength = 255
	ExampleOrderProductQuantityMin   = 1
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrder struct {
	ID          string
	Name        string
	Description string
	UserID      string
	Products    []*ExampleOrderProduct
	CreatedBy   string
	UpdatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ExampleOrderProduct struct {
	ID        string
	OrderID   string
	ProductID string
	Name      string
	Quantity  int
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrder(name, description, userID string, exampleOrderProducts []*ExampleOrderProduct, createdBy string) (*ExampleOrder, error) {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > ExampleOrderNameMaxLength {
		return nil, ExampleOrderErrInvalidName
	}

	description = strings.TrimSpace(description)
	if utf8.RuneCountInString(description) > ExampleOrderDescriptionMaxLength {
		return nil, ExampleOrderErrDescriptionTooLong
	}

	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil, ExampleOrderErrInvalidUserID
	}

	if len(exampleOrderProducts) == 0 {
		return nil, ExampleOrderErrNoProducts
	}

	exampleOrderValidatedProducts := make([]*ExampleOrderProduct, len(exampleOrderProducts))
	for i, exampleOrderProduct := range exampleOrderProducts {
		validated, err := NewExampleOrderProduct(exampleOrderProduct.ProductID, exampleOrderProduct.Name, exampleOrderProduct.Quantity, createdBy)
		if err != nil {
			return nil, err
		}
		exampleOrderValidatedProducts[i] = validated
	}

	now := time.Now().UTC()

	return &ExampleOrder{
		Name:        name,
		Description: description,
		UserID:      userID,
		Products:    exampleOrderValidatedProducts,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func NewExampleOrderProduct(productID, name string, quantity int, createdBy string) (*ExampleOrderProduct, error) {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > ExampleOrderProductNameMaxLength {
		return nil, ExampleOrderErrInvalidProductName
	}

	productID = strings.TrimSpace(productID)
	if productID == "" {
		return nil, ExampleOrderErrInvalidProductID
	}

	if quantity < ExampleOrderProductQuantityMin {
		return nil, ExampleOrderErrInvalidProductQuantity
	}

	now := time.Now().UTC()

	return &ExampleOrderProduct{
		ProductID: productID,
		Name:      name,
		Quantity:  quantity,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
