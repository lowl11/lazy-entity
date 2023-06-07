package condition_service

import "github.com/lowl11/lazy-entity/field_type"

const (
	template     = "WHERE {{CONDITION_LIST}}"
	itemTemplate = "{{CONDITION_NAME}} {{CONDITION_SIGN}} {{CONDITION_VALUE}}"
)

func getValue(valueType, value, field string) string {
	switch valueType {
	case field_type.Numeric:
		return value
	case field_type.Boolean:
		return value
	case field_type.Variable:
		return ":" + field
	case field_type.Join:
		return value
	}
	return "'" + value + "'"
}
