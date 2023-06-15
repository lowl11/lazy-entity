package update_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/internal/query_helpers/update_helper"
	"strings"
)

func (builder *Builder) Build() string {
	queries := make([]string, 0, 3)

	// main template
	main := "UPDATE " + builder.tableName
	queries = append(queries, main)

	// set template
	if len(builder.setValues) > 0 {
		queries = append(queries, "SET\n"+strings.Join(builder.setValues, ",\n"))
	} else if len(builder.setFields) > 0 {
		setList := make([]string, 0, len(builder.setFields))
		for _, item := range builder.setFields {
			setList = append(setList, update_helper.VariableField(item))
		}
		queries = append(queries, "SET\n"+strings.Join(setList, ",\n"))
	}

	// where template
	if builder.conditions.Len() > 0 {
		where := "WHERE " + builder.conditions.String()
		queries = append(queries, where)
	}

	return strings.Join(queries, "\n")
}

func (builder *Builder) Set(field string, value any) *Builder {
	builder.setValues = append(builder.setValues, "\t"+field+" = "+type_helper.ToString(value))
	return builder
}

func (builder *Builder) SetByFields(fields ...string) *Builder {
	builder.setFields = append(builder.setFields, fields...)
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
