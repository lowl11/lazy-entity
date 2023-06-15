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
