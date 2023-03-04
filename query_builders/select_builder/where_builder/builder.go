package where_builder

import "github.com/lowl11/lazy-entity/entity_models"

type Builder struct {
	conditions []entity_models.WhereCondition
	or         bool
}

func New() *Builder {
	return &Builder{
		conditions: make([]entity_models.WhereCondition, 0),
	}
}
