package application

import (
	context "context"

	exampleUserDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleUser/domain"
)

// ============================================================================
// Types
// ============================================================================

type ExampleUserUsecaseCreateMultiple struct {
	postgresqlTransaction  exampleUserDomain.ExampleUserPostgresqlTransaction
	createMultiplePostgres exampleUserDomain.ExampleUserPostgresqlCreateMultiple
}

// ============================================================================
// Constructors
// ============================================================================

func NewExampleUserUsecaseCreateMultiple(
	postgresqlTransaction exampleUserDomain.ExampleUserPostgresqlTransaction,
	createMultiplePostgres exampleUserDomain.ExampleUserPostgresqlCreateMultiple,
) exampleUserDomain.ExampleUserUsecaseCreateMultiple {
	return &ExampleUserUsecaseCreateMultiple{
		postgresqlTransaction:  postgresqlTransaction,
		createMultiplePostgres: createMultiplePostgres,
	}
}

// ============================================================================
// Methods
// ============================================================================

func (uc *ExampleUserUsecaseCreateMultiple) Execute(ctx context.Context, examples []exampleUserDomain.ExampleUser) ([]*exampleUserDomain.ExampleUser, error) {
	var createdExampleUsers []*exampleUserDomain.ExampleUser

	err := uc.postgresqlTransaction.WithinTransaction(ctx, func(ctx context.Context) error {
		createdExampleUsers = make([]*exampleUserDomain.ExampleUser, len(examples))
		for i := range examples {
			example, err := exampleUserDomain.NewExampleUser(examples[i].Name, examples[i].Description, examples[i].CreatedBy)
			if err != nil {
				return err
			}
			createdExampleUsers[i] = example
		}

		if err := uc.createMultiplePostgres.Execute(ctx, createdExampleUsers); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdExampleUsers, nil
}
