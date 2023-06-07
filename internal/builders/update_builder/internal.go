package update_builder

import "github.com/lowl11/lazy-entity/field_type"

const (
	mainTemplate = "UPDATE {{TABLE_NAME}}\nSET {{SET_TEMPLATE}}"
	setTemplate  = "{{SET_FIELD}} = {{SET_VALUE}}"
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
	case field_type.CountVariable:
		return value
	}
	return "'" + value + "'"
}
