package delete_builder

import (
	"github.com/lowl11/lazy-entity/field_type"
	"github.com/lowl11/lazy-entity/internal/entity_domain"
	"github.com/lowl11/lazy-entity/internal/signs"
	"github.com/lowl11/lazy-entity/internal/template_service"
	"strings"
)

func (builder *Builder) Build() string {
	templateList := make([]string, 0, 1)

	// main template
	main := template_service.New(mainTemplate).
		Var("TABLE_NAME", builder.tableName).
		Get()

	templateList = append(templateList, main)

	// condition template
	if len(builder.conditionList) > 0 {
		conditionService := template_service.New(conditionTemplate)

		for _, item := range builder.conditionList {
			conditionService.Var("CONDITION_NAME", item.Field)
			conditionService.Var("CONDITION_SIGN", item.Sign)
			conditionService.Var("CONDITION_VALUE", getValue(item.ValueType, item.Value))
		}

		templateList = append(templateList, conditionService.Get())
	}

	return strings.Join(templateList, "\n")
}

func (builder *Builder) ConditionEquals(field, value, valueType string) *Builder {
	builder.conditionList = append(builder.conditionList, entity_domain.ConditionPair{
		Field:     field,
		Value:     value,
		Sign:      signs.Equals,
		ValueType: valueType,
	})
	return builder
}

func (builder *Builder) ConditionLike(field, value string) *Builder {
	builder.conditionList = append(builder.conditionList, entity_domain.ConditionPair{
		Field:     field,
		Value:     value,
		Sign:      signs.Like,
		ValueType: field_type.Text,
	})
	return builder
}

func (builder *Builder) ConditionIlike(field, value string) *Builder {
	builder.conditionList = append(builder.conditionList, entity_domain.ConditionPair{
		Field:     field,
		Value:     value,
		Sign:      signs.Ilike,
		ValueType: field_type.Text,
	})
	return builder
}
