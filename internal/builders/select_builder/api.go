package select_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/join_types"
	"strings"
)

func (builder *Builder) Build() string {
	queries := make([]string, 0, 3)

	// main template
	main := "SELECT " + builder.getFields() + "\nFROM " + builder.getTableName()
	queries = append(queries, main)

	// join template
	if len(builder.joinList) > 0 {
		joinQueries := make([]string, 0, len(builder.joinList))
		for _, item := range builder.joinList {
			joinQueries = append(joinQueries, "\t"+item.joinType+" JOIN "+item.TableName+" AS "+item.AliasName+" ON "+item.Conditions)
		}
		queries = append(queries, strings.Join(joinQueries, "\n"))
	}

	// where template
	if len(builder.conditions) > 0 {
		where := "WHERE " + builder.conditions
		queries = append(queries, where)
	}

	// order by template
	if len(builder.orderFields) > 0 {
		orderQueries := make([]string, 0, len(builder.orderFields))
		for _, item := range builder.orderFields {
			orderQueries = append(orderQueries, builder.getFieldItem(item))
		}
		queries = append(queries, "ORDER BY "+strings.Join(orderQueries, ", ")+" "+builder.orderType)
	}

	// group by template
	if len(builder.groupByFields) > 0 {
		groupQueries := make([]string, 0, len(builder.groupByFields))
		for _, item := range builder.groupByFields {
			groupQueries = append(groupQueries, builder.getFieldItem(item))
		}
		queries = append(queries, "GROUP BY "+strings.Join(groupQueries, ", "))
	}

	// having template
	if builder.havingExpression != "" {
		queries = append(queries, "HAVING "+builder.havingExpression)
	}

	// offset template
	offset := builder.getOffset()
	if offset != "" {
		queries = append(queries, offset)
	}

	// limit template
	limit := builder.getLimit()
	if limit != "" {
		queries = append(queries, limit)
	}

	return strings.Join(queries, "\n")
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

func (builder *Builder) Join(tableName, aliasName, conditions string) *Builder {
	builder.joinList = append(builder.joinList, joinModel{
		TableName:  tableName,
		AliasName:  aliasName,
		Conditions: conditions,
		joinType:   join_types.Inner,
	})
	return builder
}

func (builder *Builder) LeftJoin(tableName, aliasName, conditions string) *Builder {
	builder.joinList = append(builder.joinList, joinModel{
		TableName:  tableName,
		AliasName:  aliasName,
		Conditions: conditions,
		joinType:   join_types.Left,
	})
	return builder
}

func (builder *Builder) RightJoin(tableName, aliasName, conditions string) *Builder {
	builder.joinList = append(builder.joinList, joinModel{
		TableName:  tableName,
		AliasName:  aliasName,
		Conditions: conditions,
		joinType:   join_types.Right,
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

func (builder *Builder) OrderBy(orderType string, fieldList ...string) *Builder {
	builder.orderType = orderType
	builder.orderFields = fieldList
	return builder
}

func (builder *Builder) Having(expression string) *Builder {
	builder.havingExpression = expression
	return builder
}

func (builder *Builder) GroupBy(fields ...string) *Builder {
	builder.groupByFields = append(builder.groupByFields, fields...)
	return builder
}

func (builder *Builder) Count(field string, value any, expression func(field string, value any) string) string {
	fieldName := builder.getFieldItem(field)
	expressionValue := expression(field, value)
	expressionValue = strings.ReplaceAll(expressionValue, fieldName, "")
	return "COUNT(" + fieldName + ")" + expressionValue
}

func (builder *Builder) Min(field string, value any, expression func(field string, value any) string) string {
	fieldName := builder.getFieldItem(field)
	expressionValue := expression(field, value)
	expressionValue = strings.ReplaceAll(expressionValue, fieldName, "")
	return "MIN(" + fieldName + ")" + expressionValue
}

func (builder *Builder) Max(field string, value any, expression func(field string, value any) string) string {
	fieldName := builder.getFieldItem(field)
	expressionValue := expression(field, value)
	expressionValue = strings.ReplaceAll(expressionValue, fieldName, "")
	return "MAX(" + fieldName + ")" + expressionValue
}

func (builder *Builder) Avg(field string, value any, expression func(field string, value any) string) string {
	fieldName := builder.getFieldItem(field)
	expressionValue := expression(field, value)
	expressionValue = strings.ReplaceAll(expressionValue, fieldName, "")
	return "AVG(" + fieldName + ")" + expressionValue
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

func (builder *Builder) NotEqual(field string, value any) string {
	return builder.getFieldItem(field) + " != " + type_helper.ToString(value)
}

func (builder *Builder) In(field string, value any) string {
	return builder.getFieldItem(field) + " IN " + type_helper.ToString(value)
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

func (builder *Builder) Offset(value int) *Builder {
	builder.offset = value
	return builder
}

func (builder *Builder) Limit(value int) *Builder {
	builder.limit = value
	return builder
}
