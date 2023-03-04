package type_service

import "strconv"

func GetString(value any) string {
	if isString(value) {
		return value.(string)
	}

	if isBool(value) {
		return value.(string)
	}

	if isInt(value) {
		return strconv.Itoa(value.(int))
	}

	return ""
}
