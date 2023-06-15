package update_helper

import (
	"fmt"
	"strings"
)

func VariableField(name string) string {
	field := strings.Builder{}
	field.WriteString("\t")
	field.WriteString(name)
	field.WriteString(" = :")
	field.WriteString(name)
	fmt.Println("vfl:", field.Len())
	return field.String()
}
