package insert_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"strings"
)

func (builder *Builder) getFields() string {
	return " (" + strings.Join(builder.fieldList, ", ") + ")"
}

func (builder *Builder) getVariableFields() string {
	variableList := make([]string, 0, len(builder.fieldList))
	for _, item := range builder.fieldList {
		variableList = append(variableList, ":"+item)
	}
	return strings.Join(variableList, ", ")
}

func (builder *Builder) getVariables() string {
	list := make([]string, 0, len(builder.variableList))
	for _, item := range builder.variableList {
		list = append(list, type_helper.ToString(item))
	}
	return strings.Join(list, ", ")
}
