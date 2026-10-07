package domain

import (
	time "time"
)

// ============================================================================
// Constants
// ============================================================================

const EventTypeParentCreated = "PARENT_CREATED"

// ============================================================================
// Types
// ============================================================================

type Event struct {
	Type       string
	ParentID   string
	OccurredAt time.Time
}

// ============================================================================
// Constructors
// ============================================================================

func NewEvent(eventType, parentID string) Event {
	return Event{
		Type:       eventType,
		ParentID:   parentID,
		OccurredAt: time.Now().UTC(),
	}
}
