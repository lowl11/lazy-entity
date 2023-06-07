package query_service

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/entity_domain"
	"github.com/lowl11/lazy-entity/field_type"
	"github.com/lowl11/lazy-entity/signs"
)

func (service *Service) Delete() string {
	return delete_builder.
		New(service.tableName).
		Build(service.conditionList)
}

func (service *Service) ConditionEquals(field, value, valueType string) *Service {
	service.conditionList = append(service.conditionList, entity_domain.ConditionPair{
		Field:     field,
		Value:     value,
		Sign:      signs.Equals,
		ValueType: valueType,
	})
	return service
}

func (service *Service) ConditionLike(field, value string) *Service {
	service.conditionList = append(service.conditionList, entity_domain.ConditionPair{
		Field:     field,
		Value:     value,
		Sign:      signs.Like,
		ValueType: field_type.Text,
	})
	return service
}

func (service *Service) ConditionIlike(field, value string) *Service {
	service.conditionList = append(service.conditionList, entity_domain.ConditionPair{
		Field:     field,
		Value:     value,
		Sign:      signs.Ilike,
		ValueType: field_type.Text,
	})
	return service
}
