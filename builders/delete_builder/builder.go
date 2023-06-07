package delete_builder

import "github.com/lowl11/lazy-entity/entity_domain"

type Builder struct {
	tableName     string
	conditionList []entity_domain.ConditionPair
}

func New(tableName string) *Builder {
	return &Builder{
		tableName:     tableName,
		conditionList: make([]entity_domain.ConditionPair, 0),
	}
}
