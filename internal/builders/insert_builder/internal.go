package insert_builder

import "strings"

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
