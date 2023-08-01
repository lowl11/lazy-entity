package condition_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/string_helper"
	"strings"
)

func (builder *Builder) And(conditions ...string) string {
	conditionArray := make([]string, 0, len(conditions))
	for _, item := range conditions {
		conditionArray = append(conditionArray, item)
	}
	return string_helper.Concat("(", strings.Join(conditionArray, " AND "), ")")
}

func (builder *Builder) Or(conditions ...string) string {
	conditionArray := make([]string, 0, len(conditions))
	for _, item := range conditions {
		conditionArray = append(conditionArray, item)
	}
	return string_helper.Concat("(", strings.Join(conditionArray, " OR "), ")")
}

func (builder *Builder) Equal(field string, value any) string {
	return builder.statement(field, " = ", value)
}

func (builder *Builder) NotEqual(field string, value any) string {
	return builder.statement(field, " != ", value)
}

func (builder *Builder) Is(field string, value any) string {
	return builder.statement(field, " IS ", value)
}

func (builder *Builder) NotIs(field string, value any) string {
	return builder.statement(field, " IS NOT ", value)
}

func (builder *Builder) In(field string, value any) string {
	return builder.statement(field, " = ANY(", value)
}

func (builder *Builder) NotIn(field string, value any) string {
	return builder.statement(field, " NOT IN ", value)
}

func (builder *Builder) Like(field string, value string) string {
	return builder.statement(field, " LIKE ", value)
}

func (builder *Builder) NotLike(field string, value string) string {
	return builder.statement(field, " NOT LIKE ", value)
}

func (builder *Builder) ILike(field string, value string) string {
	return builder.statement(field, " ILIKE ", value)
}

func (builder *Builder) NotILike(field string, value string) string {
	return builder.statement(field, " NOT ILIKE ", value)
}

func (builder *Builder) Gte(field string, value any) string {
	return builder.statement(field, " >= ", value)
}

func (builder *Builder) Gt(field string, value any) string {
	return builder.statement(field, " > ", value)
}

func (builder *Builder) Lte(field string, value any) string {
	return builder.statement(field, " <= ", value)
}

func (builder *Builder) Lt(field string, value any) string {
	return builder.statement(field, " < ", value)
}

func (builder *Builder) IsNull(field string) string {
	return builder.statement(field, " IS NULL", nil)
}

func (builder *Builder) IsNotNull(field string) string {
	return builder.statement(field, " IS NOT NULL", nil)
}

func (builder *Builder) Between(field string) string {
	return builder.getFieldItem(field)
}
