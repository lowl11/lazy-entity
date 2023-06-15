package update_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/internal/query_helpers/update_helper"
	"strings"
)

func (builder *Builder) Build() string {
	// main template
	query := strings.Builder{}
	query.Grow(200)
	query.WriteString("UPDATE ")
	query.WriteString(builder.tableName)
	query.WriteString("\n")

	// set template
	query.WriteString("SET\n")
	if len(builder.setValues) > 0 {
		query.WriteString(strings.Join(builder.setValues, ",\n"))
	} else if len(builder.setFields) > 0 {
		setList := make([]string, 0, len(builder.setFields))
		for _, item := range builder.setFields {
			setList = append(setList, update_helper.VariableField(item))
		}
		query.WriteString(strings.Join(setList, ",\n"))
	}
	query.WriteString("\n")

	// where template
	if builder.conditions.Len() > 0 {
		query.WriteString("WHERE ")
		query.WriteString(builder.conditions.String())
		query.WriteString("\n")
	}

	return query.String()
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
