package domain

import (
	time "time"
)

// ============================================================================
// Constants
// ============================================================================

const EventTypeOrderCreated = "ORDER_CREATED"

// ============================================================================
// Types
// ============================================================================

type ExampleOrderEvent struct {
	Type       string
	OrderID    string
	OccurredAt time.Time
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderEvent(eventType, orderID string) ExampleOrderEvent {
	return ExampleOrderEvent{
		Type:       eventType,
		OrderID:    orderID,
		OccurredAt: time.Now().UTC(),
	}
}
