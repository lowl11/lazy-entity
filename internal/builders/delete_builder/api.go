package delete_builder

import (
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
