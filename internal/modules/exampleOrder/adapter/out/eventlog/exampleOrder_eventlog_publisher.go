package eventlog

import (
	context "context"
	log "log"
	time "time"

	exampleOrderDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleOrder/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleOrderEventLogPublisher struct{}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleOrderEventLogPublisher() *ExampleOrderEventLogPublisher {
	return &ExampleOrderEventLogPublisher{}
}

// ============================================================================
// Methods
// ============================================================================

func (p *ExampleOrderEventLogPublisher) Execute(ctx context.Context, exampleOrderEvent exampleOrderDomain.Event) error {
	log.Printf("example order event published: type=%s order_id=%s occurred_at=%s", exampleOrderEvent.Type, exampleOrderEvent.OrderID, exampleOrderEvent.OccurredAt.Format(time.RFC3339))
	return nil
}
