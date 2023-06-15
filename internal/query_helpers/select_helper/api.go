package select_helper

import (
	"strings"
)

func Main(tableName, fields string) string {
	main := strings.Builder{}
	main.Grow(500)
	main.WriteString("SELECT ")
	main.WriteString(fields)
	main.WriteString("\nFROM ")
	main.WriteString(tableName)
	return main.String()
}

func OrderBy(orderType, orderQueries string) string {
	orderBy := strings.Builder{}
	orderBy.Grow(30)
	orderBy.WriteString("ORDER BY ")
	orderBy.WriteString(orderQueries)
	orderBy.WriteString(" ")
	orderBy.WriteString(orderType)
	return orderBy.String()
}

func Join(joinType, tableName, aliasName, condition string) string {
	join := strings.Builder{}
	join.Grow(100)
	join.WriteString("\t")
	join.WriteString(joinType)
	join.WriteString(" JOIN ")
	join.WriteString(tableName)
	join.WriteString(" AS ")
	join.WriteString(aliasName)
	join.WriteString(" ON (")
	join.WriteString(condition)
	join.WriteString(")")
	return join.String()
}

func Count(fieldName, expressionValue string) string {
	count := strings.Builder{}
	count.Grow(50)
	count.WriteString("COUNT(")
	count.WriteString(fieldName)
	count.WriteString(")")
	count.WriteString(expressionValue)
	return count.String()
}

func Min(fieldName, expressionValue string) string {
	min := strings.Builder{}
	min.Grow(50)
	min.WriteString("MIN(")
	min.WriteString(fieldName)
	min.WriteString(")")
	min.WriteString(expressionValue)
	return min.String()
}

func Max(fieldName, expressionValue string) string {
	max := strings.Builder{}
	max.Grow(50)
	max.WriteString("MAX(")
	max.WriteString(fieldName)
	max.WriteString(")")
	max.WriteString(expressionValue)
	return max.String()
}

func Avg(fieldName, expressionValue string) string {
	avg := strings.Builder{}
	avg.Grow(50)
	avg.WriteString("AVG(")
	avg.WriteString(fieldName)
	avg.WriteString(")")
	avg.WriteString(expressionValue)
	return avg.String()
}
