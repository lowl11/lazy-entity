package delete_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"strings"
)

func (builder *Builder) Build() string {
	queries := make([]string, 0, 2)

	// main template
	main := "DELETE FROM " + builder.tableName
	queries = append(queries, main)

	// where template
	if len(builder.conditions) > 0 {
		where := "WHERE " + builder.conditions
		queries = append(queries, where)
	}

	return strings.Join(queries, "\n")
}

func (builder *Builder) Where(conditions ...string) *Builder {
	conditionArray := make([]string, 0, len(conditions))
	for _, item := range conditions {
		conditionArray = append(conditionArray, "\n\t"+item)
	}
	builder.conditions += strings.Join(conditionArray, " AND ")
	return builder
}

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

func (builder *Builder) In(field string, value any) string {
	return builder.getFieldItem(field) + " IN " + type_helper.ToString(value)
}

func (builder *Builder) Like(field string, value string) string {
	return builder.getFieldItem(field) + " LIKE " + type_helper.ToString(value)
}

func (builder *Builder) ILike(field string, value string) string {
	return builder.getFieldItem(field) + " ILIKE " + type_helper.ToString(value)
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
