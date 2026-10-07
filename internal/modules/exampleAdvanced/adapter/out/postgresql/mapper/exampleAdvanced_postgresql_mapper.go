package mapper

import (
	postgresqlModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model"
	defaultModel "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/infrastructure/postgresql/model/default"
	exampleAdvancedDomain "github.com/kelsonwinith/learn.go-hexagonal-architecture/internal/modules/exampleAdvanced/domain"
)

// ============================================================================
// Functions
// ============================================================================

func ToParentModel(parent *exampleAdvancedDomain.Parent) *postgresqlModel.ExampleAdvancedParentModel {
	return &postgresqlModel.ExampleAdvancedParentModel{
		BaseModel: defaultModel.BaseModel{
			ID:        parent.ID,
			CreatedAt: parent.CreatedAt,
			UpdatedAt: parent.UpdatedAt,
			CreatedBy: parent.CreatedBy,
			UpdatedBy: parent.UpdatedBy,
		},
		Name:        parent.Name,
		Description: parent.Description,
	}
}

func ToChildModel(child *exampleAdvancedDomain.Child) *postgresqlModel.ExampleAdvancedChildModel {
	return &postgresqlModel.ExampleAdvancedChildModel{
		BaseModel: defaultModel.BaseModel{
			ID:        child.ID,
			CreatedAt: child.CreatedAt,
			UpdatedAt: child.UpdatedAt,
			CreatedBy: child.CreatedBy,
			UpdatedBy: child.UpdatedBy,
		},
		ParentID: child.ParentID,
		Name:     child.Name,
		Quantity: child.Quantity,
	}
}

func ToChildModels(children []*exampleAdvancedDomain.Child) []*postgresqlModel.ExampleAdvancedChildModel {
	entities := make([]*postgresqlModel.ExampleAdvancedChildModel, len(children))
	for i, child := range children {
		entities[i] = ToChildModel(child)
	}

	return entities
}

func ToParentDomain(entity *postgresqlModel.ExampleAdvancedParentModel) *exampleAdvancedDomain.Parent {
	parent := &exampleAdvancedDomain.Parent{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedBy:   entity.CreatedBy,
		UpdatedBy:   entity.UpdatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}

	if entity.Children != nil {
		parent.Children = make([]*exampleAdvancedDomain.Child, len(entity.Children))
		for i, child := range entity.Children {
			parent.Children[i] = ToChildDomain(child)
		}
	}

	return parent
}

func ToChildDomain(entity *postgresqlModel.ExampleAdvancedChildModel) *exampleAdvancedDomain.Child {
	return &exampleAdvancedDomain.Child{
		ID:        entity.ID,
		ParentID:  entity.ParentID,
		Name:      entity.Name,
		Quantity:  entity.Quantity,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
