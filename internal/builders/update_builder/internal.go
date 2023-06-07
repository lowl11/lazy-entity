package update_builder

import "github.com/lowl11/lazy-entity/field_type"

const (
	mainTemplate      = "UPDATE {{TABLE_NAME}}\nSET {{SET_TEMPLATE}}"
	setTemplate       = "{{SET_FIELD}} = {{SET_VALUE}}"
	conditionTemplate = "WHERE {{CONDITION_NAME}} {{CONDITION_SIGN}} {{CONDITION_VALUE}}"
)

func getValue(valueType, value string) string {
	switch valueType {
	case field_type.Numeric:
		return value
	case field_type.Boolean:
		return value
	}
	return "'" + value + "'"
}
