package select_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/sql_helper"
	"github.com/lowl11/lazy-entity/internal/query_helpers/select_helper"
	"github.com/lowl11/lazy-entity/join_types"
	"strings"
)

func (builder *Builder) Build() string {
	query := strings.Builder{}
	//query.Grow() // todo: implement me

	// main template
	select_helper.Main(&query, builder.getTableName(), builder.getFields())

	// join template
	for _, item := range builder.joinList {
		select_helper.Join(
			&query,
			item.joinType,
			item.TableName,
			sql_helper.AliasName(item.AliasName),
			sql_helper.ConditionAlias(item.AliasName, item.Conditions),
		)
	}

	// where template
	if builder.conditions.Len() > 0 {
		query.WriteString("WHERE ")
		query.WriteString(builder.conditions.String())
		query.WriteString("\n")
	}

	// order by template
	if len(builder.orderFields) > 0 {
		orderQueries := make([]string, 0, len(builder.orderFields))
		for _, item := range builder.orderFields {
			orderQueries = append(orderQueries, builder.getFieldItem(item))
		}

		select_helper.OrderBy(&query, builder.orderType, strings.Join(orderQueries, ", "))
	}

	// group by template
	if len(builder.groupByFields) > 0 {
		groupQueries := make([]string, 0, len(builder.groupByFields))
		for _, item := range builder.groupByFields {
			groupQueries = append(groupQueries, builder.getFieldItem(item))
		}

		select_helper.GroupBy(&query, strings.Join(groupQueries, ", "))
	}

	// having template
	select_helper.Having(&query, builder.havingExpression)

	// offset template
	if offset := builder.getOffset(); offset != "" {
		query.WriteString(offset)
		query.WriteString("\n")
	}

	// limit template
	if limit := builder.getLimit(); limit != "" {
		query.WriteString(limit)
		query.WriteString("\n")
	}

	return query.String()
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
	return select_helper.Count(fieldName, expressionValue)
}

func (builder *Builder) Min(field string, value any, expression func(field string, value any) string) string {
	fieldName := builder.getFieldItem(field)
	expressionValue := expression(field, value)
	expressionValue = strings.ReplaceAll(expressionValue, fieldName, "")
	return select_helper.Min(field, expressionValue)
}

func (builder *Builder) Max(field string, value any, expression func(field string, value any) string) string {
	fieldName := builder.getFieldItem(field)
	expressionValue := expression(field, value)
	expressionValue = strings.ReplaceAll(expressionValue, fieldName, "")
	return select_helper.Max(field, expressionValue)
}

func (builder *Builder) Avg(field string, value any, expression func(field string, value any) string) string {
	fieldName := builder.getFieldItem(field)
	expressionValue := expression(field, value)
	expressionValue = strings.ReplaceAll(expressionValue, fieldName, "")
	return select_helper.Avg(field, expressionValue)
}

func (builder *Builder) Offset(value int) *Builder {
	builder.offset = value
	return builder
}

func (builder *Builder) Limit(value int) *Builder {
	builder.limit = value
	return builder
}
