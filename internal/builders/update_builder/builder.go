package update_builder

import "github.com/lowl11/lazy-entity/internal/entity_domain"

type Builder struct {
	tableName     string
	conditionList []entity_domain.ConditionPair
	updateList    []entity_domain.UpdatePair
}

func New(tableName string) *Builder {
	return &Builder{
		tableName:     tableName,
		conditionList: make([]entity_domain.ConditionPair, 0),
		updateList:    make([]entity_domain.UpdatePair, 0),
	}
}
