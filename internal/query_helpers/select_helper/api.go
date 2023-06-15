package select_helper

import "strings"

func Main(tableName, fields string) string {
	main := strings.Builder{}
	main.WriteString("SELECT ")
	main.WriteString(fields)
	main.WriteString("\nFROM ")
	main.WriteString(tableName)
	return main.String()
}
