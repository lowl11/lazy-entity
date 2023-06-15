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
