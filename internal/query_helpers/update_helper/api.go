package update_helper

import (
	"strings"
)

func VariableField(name string) string {
	field := strings.Builder{}
	field.WriteString("\t")
	field.WriteString(name)
	field.WriteString(" = :")
	field.WriteString(name)
	return field.String()
}
