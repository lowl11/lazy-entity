package select_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/sql_helper"
	"github.com/lowl11/lazy-entity/internal/query_helpers/select_helper"
	"github.com/lowl11/lazy-entity/join_types"
	"strings"
)

func (builder *Builder) Build() string {
	queries := make([]string, 0, 10)

	// main template
	queries = append(queries, select_helper.Main(builder.getTableName(), builder.getFields()))

	// join template
	if len(builder.joinList) > 0 {
		joinQueries := make([]string, 0, len(builder.joinList))
		for _, item := range builder.joinList {
			joinQueries = append(joinQueries, select_helper.Join(
				item.joinType,
				item.TableName,
				sql_helper.AliasName(item.AliasName),
				sql_helper.ConditionAlias(item.AliasName, item.Conditions),
			))
		}
		queries = append(queries, strings.Join(joinQueries, "\n"))
	}

	// where template
	if builder.conditions.Len() > 0 {
		where := "WHERE " + builder.conditions.String()
		queries = append(queries, where)
	}

	// order by template
	if len(builder.orderFields) > 0 {
		orderQueries := make([]string, 0, len(builder.orderFields))
		for _, item := range builder.orderFields {
			orderQueries = append(orderQueries, builder.getFieldItem(item))
		}
		queries = append(queries, select_helper.OrderBy(builder.orderType, strings.Join(orderQueries, ", ")))
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
	builder.fieldList = append(builder.fieldList, fieldList...)
	return builder
}

func (builder *Builder) From(tableName string) *Builder {
	builder.tableName = tableName
	return builder
}

func (builder *Builder) Alias(aliasName string) *Builder {
	builder.aliasName = sql_helper.AliasName(aliasName)
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

	if builder.conditions.Len() == 0 {
		builder.conditions.WriteString(strings.Join(conditionArray, " AND "))
	} else {
		builder.conditions.WriteString(" AND " + strings.Join(conditionArray, " AND "))
	}

	return builder
}

func (builder *Builder) WhereOr(conditions ...string) *Builder {
	conditionArray := make([]string, 0, len(conditions))
	for _, item := range conditions {
		conditionArray = append(conditionArray, "\n\t"+item)
	}
	if builder.conditions.Len() == 0 {
		builder.conditions.WriteString(strings.Join(conditionArray, " AND "))
	} else {
		builder.conditions.WriteString(" OR " + strings.Join(conditionArray, " AND "))
	}

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

func (builder *Builder) Offset(value int) *Builder {
	builder.offset = value
	return builder
}

func (builder *Builder) Limit(value int) *Builder {
	builder.limit = value
	return builder
}
