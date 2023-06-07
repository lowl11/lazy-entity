package insert_builder

import "github.com/lowl11/lazy-entity/field_type"

const (
	mainTemplate  = "INSERT INTO {{TABLE_NAME}} ({{FIELD_LIST}})"
	valueTemplate = "values ({{VALUE_LIST}})"
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
