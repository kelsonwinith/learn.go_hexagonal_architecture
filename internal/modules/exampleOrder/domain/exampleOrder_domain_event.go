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

type Event struct {
	Type       string
	OrderID    string
	OccurredAt time.Time
}

// ============================================================================
// Constructors
// ============================================================================

func NewEvent(eventType, orderID string) Event {
	return Event{
		Type:       eventType,
		OrderID:    orderID,
		OccurredAt: time.Now().UTC(),
	}
}
