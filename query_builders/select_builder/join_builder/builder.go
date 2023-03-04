package join_builder

import "github.com/lowl11/lazy-entity/entity_models"

type Builder struct {
	tableName string
	asName    string

	conditions []entity_models.JoinCondition

	left bool
}

func New(tableName string) *Builder {
	return &Builder{
		tableName:  tableName,
		conditions: []entity_models.JoinCondition{},
	}
}
