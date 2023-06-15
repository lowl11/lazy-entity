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
	if builder.conditions.Len() > 0 {
		where := "WHERE " + builder.conditions.String()
		queries = append(queries, where)
	}

	return strings.Join(queries, "\n")
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
