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

func NewExampleOrder(name, description, userID string, products []*ExampleOrderProduct, createdBy string) (*ExampleOrder, error) {
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

	if len(products) == 0 {
		return nil, ExampleOrderErrNoProducts
	}

	validatedProducts := make([]*ExampleOrderProduct, len(products))
	for i, orderProduct := range products {
		validated, err := NewExampleOrderProduct(orderProduct.ProductID, orderProduct.Name, orderProduct.Quantity, createdBy)
		if err != nil {
			return nil, err
		}
		validatedProducts[i] = validated
	}

	now := time.Now().UTC()

	return &ExampleOrder{
		Name:        name,
		Description: description,
		UserID:      userID,
		Products:    validatedProducts,
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

// ============================================================================
// Methods
// ============================================================================

func (o *ExampleOrder) IsOwnedBy(userID string) bool {
	return o.CreatedBy == userID
}
