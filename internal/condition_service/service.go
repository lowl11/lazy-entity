package condition_service

import "github.com/lowl11/lazy-entity/internal/entity_domain"

type Service struct {
	predicate     string
	conditionList []entity_domain.ConditionPair

	aliasName string
}

func New(predicate string, conditionList []entity_domain.ConditionPair) *Service {
	return &Service{
		predicate:     predicate,
		conditionList: conditionList,
	}
}
