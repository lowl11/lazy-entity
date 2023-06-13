package condition_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"strings"
)

func (builder *Builder) And(conditions ...string) string {
	conditionArray := make([]string, 0, len(conditions))
	for _, item := range conditions {
		conditionArray = append(conditionArray, item)
	}
	return "(" + strings.Join(conditionArray, " AND ") + ")"
}

func (builder *Builder) Or(conditions ...string) string {
	conditionArray := make([]string, 0, len(conditions))
	for _, item := range conditions {
		conditionArray = append(conditionArray, item)
	}
	return "(" + strings.Join(conditionArray, " OR ") + ")"
}

func (builder *Builder) Equal(field string, value any) string {
	return builder.getFieldItem(field) + " = " + type_helper.ToString(value)
}

func (builder *Builder) NotEqual(field string, value any) string {
	return builder.getFieldItem(field) + " != " + type_helper.ToString(value)
}

func (builder *Builder) In(field string, value any) string {
	return builder.getFieldItem(field) + " = ANY(" + type_helper.ToString(value) + ")"
}

func (builder *Builder) NotIn(field string, value any) string {
	return builder.getFieldItem(field) + " NOT IN " + type_helper.ToString(value)
}

func (builder *Builder) Like(field string, value string) string {
	return builder.getFieldItem(field) + " LIKE " + type_helper.ToString(value)
}

func (builder *Builder) NotLike(field string, value string) string {
	return builder.getFieldItem(field) + " NOT LIKE " + type_helper.ToString(value)
}

func (builder *Builder) ILike(field string, value string) string {
	return builder.getFieldItem(field) + " ILIKE " + type_helper.ToString(value)
}

func (builder *Builder) NotILike(field string, value string) string {
	return builder.getFieldItem(field) + " NOT ILIKE " + type_helper.ToString(value)
}

func (builder *Builder) Gte(field string, value any) string {
	return builder.getFieldItem(field) + " >= " + type_helper.ToString(value)
}

func (builder *Builder) Gt(field string, value any) string {
	return builder.getFieldItem(field) + " > " + type_helper.ToString(value)
}

func (builder *Builder) Lte(field string, value any) string {
	return builder.getFieldItem(field) + " <= " + type_helper.ToString(value)
}

func (builder *Builder) Lt(field string, value any) string {
	return builder.getFieldItem(field) + " < " + type_helper.ToString(value)
}

func (builder *Builder) IsNUll(field string) string {
	return builder.getFieldItem(field) + " IS NULL"
}

func (builder *Builder) IsNotNull(field string) string {
	return builder.getFieldItem(field) + " IS NOT NULL"
}
