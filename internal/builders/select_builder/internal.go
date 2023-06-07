package select_builder

import "github.com/lowl11/lazy-entity/field_type"

const (
	mainTemplate      = "SELECT {{FIELD_LIST}} FROM {{TABLE_NAME}}{{ALIAS_NAME}}"
	conditionTemplate = "WHERE {{CONDITION_NAME}} {{CONDITION_SIGN}} {{CONDITION_VALUE}}"
)

func getValue(valueType, value, field string) string {
	switch valueType {
	case field_type.Numeric:
		return value
	case field_type.Boolean:
		return value
	case field_type.Variable:
		return ":" + field
	}
	return "'" + value + "'"
}
