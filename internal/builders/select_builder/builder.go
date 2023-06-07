package select_builder

import "github.com/lowl11/lazy-entity/internal/entity_domain"

type Builder struct {
	tableName string
	aliasName string

	fieldList     []string
	conditionList []entity_domain.ConditionPair
	joinList      []entity_domain.JoinPair
}

func New(tableName string) *Builder {
	return &Builder{
		tableName:     tableName,
		fieldList:     make([]string, 0),
		conditionList: make([]entity_domain.ConditionPair, 0),
		joinList:      make([]entity_domain.JoinPair, 0),
	}
}
