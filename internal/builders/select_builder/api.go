package select_builder

import (
	"github.com/lowl11/lazy-entity/field_type"
	"github.com/lowl11/lazy-entity/internal/entity_domain"
	"github.com/lowl11/lazy-entity/internal/services/condition_service"
	"github.com/lowl11/lazy-entity/internal/services/join_service"
	"github.com/lowl11/lazy-entity/internal/services/template_service"
	"github.com/lowl11/lazy-entity/internal/signs"
	"github.com/lowl11/lazy-entity/predicates"
	"strings"
)

func (builder *Builder) Build() string {
	templateList := make([]string, 0, 1)

	// field list
	var fieldListValue string
	for index, item := range builder.fieldList {
		var fieldText string
		if index >= len(builder.fieldList)-1 {
			fieldText = item
		} else {
			fieldText = item + ", "
		}

		if len(builder.aliasName) > 0 && !strings.Contains(fieldText, ".") {
			fieldText = builder.aliasName + "." + fieldText
		}
		fieldListValue += fieldText
	}

	// if no fields, search all (*)
	if fieldListValue == "" {
		fieldListValue = "*"
	}

	// main template
	mainService := template_service.New(mainTemplate).
		Var("TABLE_NAME", builder.tableName).
		Var("FIELD_LIST", fieldListValue)

	// set alias name
	if len(builder.aliasName) > 0 {
		mainService.Var("ALIAS_NAME", " AS "+builder.aliasName)
	} else {
		mainService.Var("ALIAS_NAME", "")
	}

	templateList = append(templateList, mainService.Get())

	// join template
	joinTemplate := join_service.New(builder.aliasName, builder.joinList).Get()
	if joinTemplate != "" {
		templateList = append(templateList, joinTemplate)
	}

	// condition template
	if len(builder.conditionList) > 0 {
		templateList = append(templateList, condition_service.New(predicates.And, builder.conditionList).
			Alias(builder.aliasName).
			Get())
	}

	return strings.Join(templateList, "\n")
}

func (builder *Builder) Alias(aliasName string) *Builder {
	builder.aliasName = aliasName
	return builder
}

func (builder *Builder) Field(fields ...string) *Builder {
	builder.fieldList = append(builder.fieldList, fields...)
	return builder
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

func (builder *Builder) Join(joins ...entity_domain.JoinPair) *Builder {
	builder.joinList = append(builder.joinList, joins...)
	return builder
}
