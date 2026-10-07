package application

import (
	context "context"

	exampleAdvancedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleAdvancedUsecaseCreate struct {
	postgresqlTransaction  exampleAdvancedDomain.ExampleAdvancedPostgresqlTransaction
	createPostgres         exampleAdvancedDomain.ExampleAdvancedPostgresqlCreate
	createChildrenPostgres exampleAdvancedDomain.ExampleAdvancedChildPostgresqlCreateMultiple
	eventPublisher         exampleAdvancedDomain.ExampleAdvancedEventPublisher
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleAdvancedUsecaseCreate(
	postgresqlTransaction exampleAdvancedDomain.ExampleAdvancedPostgresqlTransaction,
	createPostgres exampleAdvancedDomain.ExampleAdvancedPostgresqlCreate,
	createChildrenPostgres exampleAdvancedDomain.ExampleAdvancedChildPostgresqlCreateMultiple,
	eventPublisher exampleAdvancedDomain.ExampleAdvancedEventPublisher,
) exampleAdvancedDomain.ExampleAdvancedUsecaseCreate {
	return &ExampleAdvancedUsecaseCreate{
		postgresqlTransaction:  postgresqlTransaction,
		createPostgres:         createPostgres,
		createChildrenPostgres: createChildrenPostgres,
		eventPublisher:         eventPublisher,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleAdvancedUsecaseCreate) Execute(ctx context.Context, input exampleAdvancedDomain.Parent) (*exampleAdvancedDomain.Parent, error) {
	parent, err := exampleAdvancedDomain.NewParent(input.Name, input.Description, input.Children, input.CreatedBy)
	if err != nil {
		return nil, err
	}

	err = uc.postgresqlTransaction.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := uc.createPostgres.Execute(ctx, parent); err != nil {
			return err
		}

		for _, child := range parent.Children {
			child.ParentID = parent.ID
		}

		return uc.createChildrenPostgres.Execute(ctx, parent.Children)
	})
	if err != nil {
		return nil, err
	}

	event := exampleAdvancedDomain.NewEvent(exampleAdvancedDomain.EventTypeParentCreated, parent.ID)
	if err := uc.eventPublisher.Execute(ctx, event); err != nil {
		return nil, err
	}

	return parent, nil
}
