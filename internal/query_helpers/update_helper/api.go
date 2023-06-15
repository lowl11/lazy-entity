package update_helper

import (
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
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

func SetValue(field string, value any) string {
	valueBuilder := strings.Builder{}
	valueBuilder.Grow(50)
	valueBuilder.WriteString("\t")
	valueBuilder.WriteString(field)
	valueBuilder.WriteString(" = ")
	valueBuilder.WriteString(type_helper.ToString(value))
	return valueBuilder.String()
}
