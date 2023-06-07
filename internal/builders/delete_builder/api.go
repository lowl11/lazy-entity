package delete_builder

import (
	"github.com/lowl11/lazy-entity/field_type"
	"github.com/lowl11/lazy-entity/internal/entity_domain"
	"github.com/lowl11/lazy-entity/internal/services/condition_service"
	"github.com/lowl11/lazy-entity/internal/services/template_service"
	"github.com/lowl11/lazy-entity/internal/signs"
	"github.com/lowl11/lazy-entity/predicates"
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
	templateList = append(templateList, condition_service.New(predicates.And, builder.conditionList).Get())

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
