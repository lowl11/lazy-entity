package condition_service

import (
	"github.com/lowl11/lazy-entity/internal/services/template_service"
	"strings"
)

func (service *Service) Get() string {
	if len(service.conditionList) == 0 {
		return ""
	}

	itemList := make([]string, 0, len(service.conditionList))
	for _, item := range service.conditionList {
		fieldText := item.Field
		if len(service.aliasName) > 0 {
			fieldText = service.aliasName + "." + fieldText
		}

		itemList = append(itemList, template_service.New(itemTemplate).
			Var("CONDITION_NAME", fieldText).
			Var("CONDITION_SIGN", item.Sign).
			Var("CONDITION_VALUE", getValue(item.ValueType, item.Value, item.Field)).
			Get(),
		)
	}

	if service.noWhere {
		return strings.Join(itemList, " "+service.predicate+" ")
	}

	return template_service.New(template).
		Var("CONDITION_LIST", strings.Join(itemList, " "+service.predicate+" ")).
		Get()
}

func (service *Service) Alias(aliasName string) *Service {
	service.aliasName = aliasName
	return service
}

func (service *Service) NoWhere() *Service {
	service.noWhere = true
	return service
}
