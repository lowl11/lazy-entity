package select_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"strings"
)

func (builder *Builder) Build() string {
	// main template
	main := "SELECT " + builder.getFields() + " FROM " + builder.getTableName()

	// where template
	var where string
	if len(builder.conditions) > 0 {
		where += "WHERE \n" + builder.conditions
	}

	return strings.Join([]string{main, where}, "\n")
}

func (builder *Builder) Fields(fieldList ...string) *Builder {
	builder.fieldList = fieldList
	return builder
}

func (builder *Builder) From(tableName string) *Builder {
	builder.tableName = tableName
	return builder
}

func (builder *Builder) Alias(aliasName string) *Builder {
	builder.aliasName = aliasName
	return builder
}

func (builder *Builder) Join(tableName, aliasName string) *Builder {
	builder.joinList = append(builder.joinList, joinModel{
		TableName: tableName,
		AliasName: aliasName,
	})
	return builder
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

func (builder *Builder) Like(field string, value any) string {
	return builder.getFieldItem(field) + " LIKE " + type_helper.ToString(value)
}

func (builder *Builder) ILike(field string, value any) string {
	return builder.getFieldItem(field) + " ILIKE " + type_helper.ToString(value)
}
