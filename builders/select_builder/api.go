package select_builder

import (
	"github.com/lowl11/lazy-entity/enums/join_types"
	"github.com/lowl11/lazy-entity/internal/grow_values"
	"github.com/lowl11/lazy-entity/internal/helpers/sql_helper"
	"github.com/lowl11/lazy-entity/internal/query_helpers/select_helper"
	"strings"
)

func (builder *Builder) Build() string {
	query := strings.Builder{}
	query.Grow(builder.grow)

	// main template
	query.WriteString("SELECT ")
	builder.getFields(&query)
	query.WriteString("\nFROM ")
	builder.getTableName(&query)
	query.WriteString("\n")

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
		query.WriteString("ORDER BY ")

		for index, item := range builder.orderFields {
			query.WriteString(builder.getFieldItem(item))

			if index < len(builder.orderFields)-1 {
				query.WriteString(", ")
			}
		}

		query.WriteString(" ")
		query.WriteString(builder.orderType)
		query.WriteString("\n")
	}

	// group by template
	if len(builder.groupByFields) > 0 {
		query.WriteString("GROUP BY ")
		for index, item := range builder.groupByFields {
			query.WriteString(builder.getFieldItem(item))

			if index < len(builder.groupByFields)-1 {
				query.WriteString(", ")
			}
		}
		query.WriteString("\n")
	}

	// having template
	select_helper.Having(&query, builder.havingExpression)

	// offset template
	builder.getOffset(&query)

	// limit template
	builder.getLimit(&query)

	return query.String()
}

func (builder *Builder) Fields(fieldList ...string) *Builder {
	builder.fieldList = append(builder.fieldList, fieldList...)

	// calculate grow
	builder.grow += grow_values.AvgFieldLen * len(fieldList)
	return builder
}

func (builder *Builder) From(tableName string) *Builder {
	builder.tableName = tableName

	// calculate grow
	builder.grow += grow_values.SelectKeyword + grow_values.FromKeyword + len(tableName)
	return builder
}

func (builder *Builder) Alias(aliasName string) *Builder {
	builder.aliasName = sql_helper.AliasName(aliasName)

	// calculate grow
	builder.grow += grow_values.AsKeyword + len(aliasName)
	return builder
}

func (builder *Builder) Join(tableName, aliasName, conditions string) *Builder {
	builder.joinList = append(builder.joinList, joinModel{
		TableName:  tableName,
		AliasName:  aliasName,
		Conditions: conditions,
		joinType:   join_types.Inner,
	})

	// calculate grow
	builder.grow += grow_values.InnerJoinKeyword + len(tableName) + len(aliasName) + len(conditions)
	return builder
}

func (builder *Builder) LeftJoin(tableName, aliasName, conditions string) *Builder {
	builder.joinList = append(builder.joinList, joinModel{
		TableName:  tableName,
		AliasName:  aliasName,
		Conditions: conditions,
		joinType:   join_types.Left,
	})

	// calculate grow
	builder.grow += grow_values.LeftJoinKeyword + len(tableName) + len(aliasName) + len(conditions)
	return builder
}

func (builder *Builder) RightJoin(tableName, aliasName, conditions string) *Builder {
	builder.joinList = append(builder.joinList, joinModel{
		TableName:  tableName,
		AliasName:  aliasName,
		Conditions: conditions,
		joinType:   join_types.Right,
	})

	// calculate grow
	builder.grow += grow_values.RightJoinKeyword + len(tableName) + len(aliasName) + len(conditions)
	return builder
}

func (builder *Builder) Where(conditions ...string) *Builder {
	conditionArray := strings.Builder{}

	if builder.conditions.Len() != 0 {
		conditionArray.WriteString(" AND ")
	}

	for _, item := range conditions {
		conditionArray.WriteString("\n\t")
		conditionArray.WriteString(item)
		conditionArray.WriteString(" AND ")
	}

	query := conditionArray.String()[:conditionArray.Len()-5]
	builder.conditions.WriteString(query)

	// calculate grow
	builder.grow += grow_values.WhereKeyword + len(query)
	return builder
}

func (builder *Builder) WhereOr(conditions ...string) *Builder {
	conditionArray := strings.Builder{}

	if builder.conditions.Len() != 0 {
		conditionArray.WriteString(" OR ")
	}

	for _, item := range conditions {
		conditionArray.WriteString("\n\t")
		conditionArray.WriteString(item)
		conditionArray.WriteString(" AND ")
	}

	query := conditionArray.String()[:conditionArray.Len()-5]
	builder.conditions.WriteString(query)

	// calculate grow
	builder.grow += grow_values.WhereKeyword + len(query)
	return builder
}

func (builder *Builder) OrderBy(orderType string, fieldList ...string) *Builder {
	builder.orderType = orderType
	builder.orderFields = fieldList

	// calculate grow
	builder.grow += grow_values.OrderByKeyword + grow_values.AvgFieldLen*len(fieldList)
	return builder
}

func (builder *Builder) Having(expression string) *Builder {
	builder.havingExpression = expression

	// calculate grow
	builder.grow += grow_values.HavingKeyword + grow_values.AvgFieldLen
	return builder
}

func (builder *Builder) GroupBy(fields ...string) *Builder {
	builder.groupByFields = append(builder.groupByFields, fields...)

	// calculate grow
	builder.grow += grow_values.GroupByKeyword + grow_values.AvgFieldLen
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

	// calculate grow
	builder.grow += grow_values.OffsetKeyword + grow_values.AvgNumLen
	return builder
}

func (builder *Builder) Limit(value int) *Builder {
	builder.limit = value

	// calculate grow
	builder.grow += grow_values.LimitKeyword + grow_values.AvgNumLen
	return builder
}
