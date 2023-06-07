package delete_builder

import (
	"github.com/lowl11/lazy-entity/entity_domain"
	"github.com/lowl11/lazy-entity/template_service"
	"strings"
)

func (builder *Builder) Build(conditionList []entity_domain.ConditionPair) string {
	templateList := make([]string, 0, 1)

	// main template
	main := template_service.
		New(mainTemplate).
		Var("TABLE_NAME", builder.tableName).
		Get()

	templateList = append(templateList, main)

	// condition template
	if len(conditionList) > 0 {
		conditionService := template_service.New(conditionTemplate)

		for _, item := range conditionList {
			conditionService.Var("CONDITION_NAME", item.Field)
			conditionService.Var("CONDITION_SIGN", item.Sign)
			conditionService.Var("CONDITION_VALUE", getValue(item.ValueType, item.Value))
		}

		templateList = append(templateList, conditionService.Get())
	}

	return strings.Join(templateList, "\n")
}
