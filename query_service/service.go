package query_service

import "github.com/lowl11/lazy-entity/entity_domain"

type Service struct {
	tableName     string
	conditionList []entity_domain.ConditionPair

	aliasName string
}

func New(tableName string) *Service {
	return &Service{
		tableName:     tableName,
		conditionList: make([]entity_domain.ConditionPair, 0),
	}
}
