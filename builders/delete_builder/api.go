package delete_builder

import (
	"strings"
)

func (builder *Builder) Build() string {
	query := strings.Builder{}

	// main template
	query.WriteString("DELETE FROM ")
	query.WriteString(builder.tableName)
	query.WriteString("\n")

	// where template
	if builder.conditions.Len() > 0 {
		query.WriteString("WHERE ")
		query.WriteString(builder.conditions.String())
		query.WriteString("\n")
	}

	return query.String()
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
