package select_helper

import (
	"strings"
)

func OrderBy(query *strings.Builder, orderType, orderQueries string) {
	query.WriteString("ORDER BY ")
	query.WriteString(orderQueries)
	query.WriteString(" ")
	query.WriteString(orderType)
	query.WriteString("\n")
}

func GroupBy(query *strings.Builder, groupQueries string) {
	query.WriteString("GROUP BY ")
	query.WriteString(groupQueries)
	query.WriteString("\n")
}

func Having(query *strings.Builder, expression string) {
	if expression == "" {
		return
	}

	query.WriteString("HAVING ")
	query.WriteString(expression)
	query.WriteString("\n")
}

func Join(query *strings.Builder, joinType, tableName, aliasName, condition string) {
	query.WriteString("\t")
	query.WriteString(joinType)
	query.WriteString(" JOIN ")
	query.WriteString(tableName)
	query.WriteString(" AS ")
	query.WriteString(aliasName)
	query.WriteString(" ON (")
	query.WriteString(condition)
	query.WriteString(")")
	query.WriteString("\n")
}

func Count(fieldName, expressionValue string) string {
	count := strings.Builder{}
	count.Grow(len(fieldName) + len(expressionValue))
	count.WriteString("COUNT(")
	count.WriteString(fieldName)
	count.WriteString(")")
	count.WriteString(expressionValue)
	return count.String()
}

func Min(fieldName, expressionValue string) string {
	min := strings.Builder{}
	min.Grow(len(fieldName) + len(expressionValue))
	min.WriteString("MIN(")
	min.WriteString(fieldName)
	min.WriteString(")")
	min.WriteString(expressionValue)
	return min.String()
}

func Max(fieldName, expressionValue string) string {
	max := strings.Builder{}
	max.Grow(len(fieldName) + len(expressionValue))
	max.WriteString("MAX(")
	max.WriteString(fieldName)
	max.WriteString(")")
	max.WriteString(expressionValue)
	return max.String()
}

func Avg(fieldName, expressionValue string) string {
	avg := strings.Builder{}
	avg.Grow(len(fieldName) + len(expressionValue))
	avg.WriteString("AVG(")
	avg.WriteString(fieldName)
	avg.WriteString(")")
	avg.WriteString(expressionValue)
	return avg.String()
}
