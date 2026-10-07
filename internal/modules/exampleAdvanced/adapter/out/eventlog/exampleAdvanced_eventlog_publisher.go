package eventlog

import (
	context "context"
	log "log"
	time "time"

	exampleAdvancedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleAdvancedEventLogPublisher struct{}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleAdvancedEventLogPublisher() *ExampleAdvancedEventLogPublisher {
	return &ExampleAdvancedEventLogPublisher{}
}

// ============================================================================
// Methods
// ============================================================================

func (p *ExampleAdvancedEventLogPublisher) Execute(ctx context.Context, event exampleAdvancedDomain.Event) error {
	log.Printf("example advanced event published: type=%s parent_id=%s occurred_at=%s", event.Type, event.ParentID, event.OccurredAt.Format(time.RFC3339))
	return nil
}
